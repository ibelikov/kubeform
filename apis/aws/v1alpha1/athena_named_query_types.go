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

type AthenaNamedQuery struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AthenaNamedQuerySpec   `json:"spec,omitempty"`
	Status            AthenaNamedQueryStatus `json:"status,omitempty"`
}

type AthenaNamedQuerySpec struct {
	Database string `json:"database"`
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	Query       string `json:"query"`
}

type AthenaNamedQueryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AthenaNamedQueryList is a list of AthenaNamedQuerys
type AthenaNamedQueryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AthenaNamedQuery CRD objects
	Items []AthenaNamedQuery `json:"items,omitempty"`
}
