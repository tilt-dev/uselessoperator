/*


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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// WebSpec defines the desired state of Web
type WebSpec struct {
	Machine string `json:"machine"`
}

// WebStatus defines the observed state of Web
type WebStatus struct {
	// +optional
	Status string `json:"status"`
}

// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="type",type="string",JSONPath=".spec.machinetype",format="byte"
// +kubebuilder:printcolumn:name="status",type="string",JSONPath=".status.status",format="byte"
// +kubebuilder:subresource:status

// Web is the Schema for the machines API
type Web struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebSpec   `json:"spec,omitempty"`
	Status WebStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebList contains a list of Web
type WebList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Web `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Web{}, &WebList{})
}
