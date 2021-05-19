// Copyright © 2021 Kaleido, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package events

import (
	"context"

	"github.com/google/uuid"
	"github.com/kaleido-io/firefly/internal/config"
	"github.com/kaleido-io/firefly/internal/events/etfactory"
	"github.com/kaleido-io/firefly/internal/i18n"
	"github.com/kaleido-io/firefly/internal/log"
	"github.com/kaleido-io/firefly/internal/retry"
	"github.com/kaleido-io/firefly/pkg/blockchain"
	"github.com/kaleido-io/firefly/pkg/database"
	"github.com/kaleido-io/firefly/pkg/publicstorage"
)

type EventManager interface {
	blockchain.Callbacks

	NewEvents() chan<- *uuid.UUID
	Start() error
	WaitStop()
}

type eventManager struct {
	ctx           context.Context
	publicstorage publicstorage.Plugin
	database      database.Plugin
	subManagers   map[string]*subscriptionManager
	retry         retry.Retry
	aggregator    *aggregator
}

func NewEventManager(ctx context.Context, pi publicstorage.Plugin, di database.Plugin) (EventManager, error) {
	if pi == nil || di == nil {
		return nil, i18n.NewError(ctx, i18n.MsgInitializationNilDepError)
	}
	em := &eventManager{
		ctx:           log.WithLogField(ctx, "role", "event-manager"),
		publicstorage: pi,
		database:      di,
		subManagers:   make(map[string]*subscriptionManager),
		retry: retry.Retry{
			InitialDelay: config.GetDuration(config.EventAggregatorRetryInitDelay),
			MaximumDelay: config.GetDuration(config.EventAggregatorRetryMaxDelay),
			Factor:       config.GetFloat64(config.EventAggregatorRetryFactor),
		},
		aggregator: newAggregator(ctx, di),
	}

	enabledTransports := config.GetStringSlice(config.EventTransportsEnabled)
	for _, transport := range enabledTransports {
		et, err := etfactory.GetPlugin(ctx, transport)
		if err != nil {
			return nil, err
		}
		em.subManagers[transport], err = newSubscriptionManager(ctx, di, et)
		if err != nil {
			return nil, err
		}
	}

	return em, nil
}

func (em *eventManager) Start() error {
	return em.aggregator.start()
}

func (em *eventManager) NewEvents() chan<- *uuid.UUID {
	return em.aggregator.eventPoller.newEvents
}

func (em *eventManager) WaitStop() {
	<-em.aggregator.eventPoller.closed
}
