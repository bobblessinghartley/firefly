// Copyright © 2022 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
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

package fftypes

type BlockchainEvent struct {
	ID         *UUID          `ffstruct:"BlockchainEvent" json:"id,omitempty"`
	Source     string         `ffstruct:"BlockchainEvent" json:"source,omitempty"`
	Namespace  string         `ffstruct:"BlockchainEvent" json:"namespace,omitempty"`
	Name       string         `ffstruct:"BlockchainEvent" json:"name,omitempty"`
	Listener   *UUID          `ffstruct:"BlockchainEvent" json:"listener,omitempty"`
	ProtocolID string         `ffstruct:"BlockchainEvent" json:"protocolId,omitempty"`
	Output     JSONObject     `ffstruct:"BlockchainEvent" json:"output,omitempty"`
	Info       JSONObject     `ffstruct:"BlockchainEvent" json:"info,omitempty"`
	Timestamp  *FFTime        `ffstruct:"BlockchainEvent" json:"timestamp,omitempty"`
	TX         TransactionRef `ffstruct:"BlockchainEvent" json:"tx"`
}
