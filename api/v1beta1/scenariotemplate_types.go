/*
Copyright 2023.

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

package v1beta1

import (
	mqttclient "github.com/paolerm/orca-mqtt-client/api/v1beta1"
	opcuaserver "github.com/paolerm/orca-opcua-server/api/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ScenarioTemplateSpec defines the desired state of ScenarioTemplate
type ScenarioTemplateSpec struct {
	OpcuaSiteSpec      []opcuaserver.OpcuaServerSpec `json:"opcuaSiteSpec"`
	MqttClientSiteSpec []mqttclient.MqttClientSpec   `json:"mqttClientSiteSpec"`
}

// ScenarioTemplateStatus defines the observed state of ScenarioTemplate
type ScenarioTemplateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ScenarioTemplate is the Schema for the scenariotemplates API
type ScenarioTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ScenarioTemplateSpec   `json:"spec,omitempty"`
	Status ScenarioTemplateStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ScenarioTemplateList contains a list of ScenarioTemplate
type ScenarioTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ScenarioTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ScenarioTemplate{}, &ScenarioTemplateList{})
}
