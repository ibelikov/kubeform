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

type LoggingProjectExclusion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoggingProjectExclusionSpec   `json:"spec,omitempty"`
	Status            LoggingProjectExclusionStatus `json:"status,omitempty"`
}

type LoggingProjectExclusionSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Disabled bool   `json:"disabled,omitempty"`
	Filter   string `json:"filter"`
	Name     string `json:"name"`
}

type LoggingProjectExclusionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LoggingProjectExclusionList is a list of LoggingProjectExclusions
type LoggingProjectExclusionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LoggingProjectExclusion CRD objects
	Items []LoggingProjectExclusion `json:"items,omitempty"`
}