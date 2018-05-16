/*
Copyright 2018 The Federation v2 Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package federatedtypes

import (
	"github.com/kubernetes-sigs/federation-v2/pkg/controller/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	serviceNamespaced bool                = true
	ServiceTypeConfig FederatedTypeConfig = FederatedTypeConfig{
		ComparisonType: util.Generation,
		Template: FederationAPIResource{
			APIResource: apiResource("FederatedService", "federatedservices", serviceNamespaced),
		},
		Placement: FederationAPIResource{
			APIResource: apiResource("FederatedServicePlacement", "federatedserviceplacements", serviceNamespaced),
		},
		Target: metav1.APIResource{
			Name:       "services",
			Group:      "",
			Kind:       "Service",
			Version:    "v1",
			Namespaced: serviceNamespaced,
		},
	}
)

func init() {
	RegisterFederatedTypeConfig(ServiceTypeConfig)
}