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

type DbSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbSnapshotSpec   `json:"spec,omitempty"`
	Status            DbSnapshotStatus `json:"status,omitempty"`
}

type DbSnapshotSpec struct {
	DbInstanceIdentifier string `json:"db_instance_identifier"`
	DbSnapshotIdentifier string `json:"db_snapshot_identifier"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DbSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DbSnapshotList is a list of DbSnapshots
type DbSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DbSnapshot CRD objects
	Items []DbSnapshot `json:"items,omitempty"`
}
