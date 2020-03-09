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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EmailConfigSpec defines the desired state of EmailConfig
type EmailConfigSpec struct {
	// The sender address.
	From string `json:"from"`
	// The address of the SMTP server.
	SmartHost string `json:"smartHost"`
	// The hostname to use when identifying to the SMTP server.
	Hello *string `json:"hello,omitempty"`
	// The username for CRAM-MD5, LOGIN and PLAIN authentications.
	AuthUsername *string `json:"authUsername,omitempty"`
	// The identity for PLAIN authentication.
	AuthIdentify *string `json:"authIdentify,omitempty"`
	// The secret contains the SMTP password for LOGIN and PLAIN authentications.
	AuthPassword *v1.SecretKeySelector `json:"authPassword,omitempty"`
	// The secret contains the SMTP secret for CRAM-MD5 authentication.
	AuthSecret *v1.SecretKeySelector `json:"authSecret,omitempty"`
}

// EmailConfigStatus defines the observed state of EmailConfig
type EmailConfigStatus struct {
}

// +kubebuilder:object:root=true

// EmailConfig is the Schema for the emailconfigs API
type EmailConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EmailConfigSpec   `json:"spec,omitempty"`
	Status EmailConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmailConfigList contains a list of EmailConfig
type EmailConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmailConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EmailConfig{}, &EmailConfigList{})
}
