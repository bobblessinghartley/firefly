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

package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly/internal/coreconfig"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/oapispec"
	"github.com/hyperledger/firefly/pkg/fftypes"
)

var postTokenApproval = &oapispec.Route{
	Name:   "postTokenApproval",
	Path:   "namespaces/{ns}/tokens/approvals",
	Method: http.MethodPost,
	PathParams: []*oapispec.PathParam{
		{Name: "ns", ExampleFromConf: coreconfig.NamespacesDefault, Description: coremsgs.APIParamsNamespace},
	},
	QueryParams: []*oapispec.QueryParam{
		{Name: "confirm", Description: coremsgs.APIConfirmQueryParam, IsBool: true},
	},
	FilterFactory:   nil,
	Description:     coremsgs.APIEndpointsPostTokenApproval,
	JSONInputValue:  func() interface{} { return &fftypes.TokenApprovalInput{} },
	JSONOutputValue: func() interface{} { return &fftypes.TokenApproval{} },
	JSONOutputCodes: []int{http.StatusAccepted, http.StatusOK},
	JSONHandler: func(r *oapispec.APIRequest) (output interface{}, err error) {
		waitConfirm := strings.EqualFold(r.QP["confirm"], "true")
		r.SuccessStatus = syncRetcode(waitConfirm)
		return getOr(r.Ctx).Assets().TokenApproval(r.Ctx, r.PP["ns"], r.Input.(*fftypes.TokenApprovalInput), waitConfirm)
	},
}
