package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type PolicyDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicyDefinitionSpec   `json:"spec,omitempty"`
	Status            PolicyDefinitionStatus `json:"status,omitempty"`
}

type PolicyDefinitionSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	DisplayName string `json:"display_name"`
	// +optional
	ManagementGroupId string `json:"management_group_id,omitempty"`
	Mode              string `json:"mode"`
	Name              string `json:"name"`
	// +optional
	Parameters string `json:"parameters,omitempty"`
	// +optional
	PolicyRule string `json:"policy_rule,omitempty"`
	PolicyType string `json:"policy_type"`
}

type PolicyDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PolicyDefinitionList is a list of PolicyDefinitions
type PolicyDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PolicyDefinition CRD objects
	Items []PolicyDefinition `json:"items,omitempty"`
}
