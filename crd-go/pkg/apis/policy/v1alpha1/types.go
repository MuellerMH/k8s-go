package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//
// HealthCheckPolicy specifies health check requirements for pods.
type HealthCheckPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HealthCheckPolicySpec   `json:"spec"`
	Status HealthCheckPolicyStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HealthCheckPolicyList is a list of HealthCheckPolicies resources
type HealthCheckPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []HealthCheckPolicy `json:"items"`
}

type HealthCheckPolicySpec struct {
}

type HealthCheckPolicyStatus struct {
	PodsFailed int64 `json:"podsFailed,omitempty"`
}
