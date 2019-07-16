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

type BatchJobQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BatchJobQueueSpec   `json:"spec,omitempty"`
	Status            BatchJobQueueStatus `json:"status,omitempty"`
}

type BatchJobQueueSpec struct {
	// +kubebuilder:validation:MaxItems=3
	ComputeEnvironments []string `json:"compute_environments"`
	Name                string   `json:"name"`
	Priority            int      `json:"priority"`
	State               string   `json:"state"`
}

type BatchJobQueueStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BatchJobQueueList is a list of BatchJobQueues
type BatchJobQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BatchJobQueue CRD objects
	Items []BatchJobQueue `json:"items,omitempty"`
}
