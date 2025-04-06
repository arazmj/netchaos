/*
Copyright 2025.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// IcmpDropRuleStatus defines the observed state of IcmpDropRule.
type IcmpDropRuleSpec struct {
	DropEvery int `json:"dropEvery"`
}

// IcmpDropRuleStatus defines the observed state of IcmpDropRule
type IcmpDropRuleStatus struct {
	// Add status fields here if needed
	Active bool `json:"active,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// IcmpDropRule is the Schema for the icmpdroprules API.
type IcmpDropRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IcmpDropRuleSpec   `json:"spec,omitempty"`
	Status IcmpDropRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IcmpDropRuleList contains a list of IcmpDropRule.
type IcmpDropRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IcmpDropRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IcmpDropRule{}, &IcmpDropRuleList{})
}
