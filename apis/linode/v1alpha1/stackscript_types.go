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

type Stackscript struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StackscriptSpec   `json:"spec,omitempty"`
	Status            StackscriptStatus `json:"status,omitempty"`
}

type StackscriptSpecUserDefinedFields struct {
	ManyOf  string `json:"many_of"`
	Default string `json:"default"`
	Label   string `json:"label"`
	Name    string `json:"name"`
	Example string `json:"example"`
	OneOf   string `json:"one_of"`
}

type StackscriptSpec struct {
	RevNote           string            `json:"rev_note"`
	DeploymentsActive int               `json:"deployments_active"`
	DeploymentsTotal  int               `json:"deployments_total"`
	Username          string            `json:"username"`
	UserDefinedFields []StackscriptSpec `json:"user_defined_fields"`
	Label             string            `json:"label"`
	Description       string            `json:"description"`
	Images            []string          `json:"images"`
	UserGravatarId    string            `json:"user_gravatar_id"`
	Created           string            `json:"created"`
	Updated           string            `json:"updated"`
	Script            string            `json:"script"`
	IsPublic          bool              `json:"is_public"`
}



type StackscriptStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StackscriptList is a list of Stackscripts
type StackscriptList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Stackscript CRD objects
	Items []Stackscript `json:"items,omitempty"`
}