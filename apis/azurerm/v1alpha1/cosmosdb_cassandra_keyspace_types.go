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

type CosmosdbCassandraKeyspace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CosmosdbCassandraKeyspaceSpec   `json:"spec,omitempty"`
	Status            CosmosdbCassandraKeyspaceStatus `json:"status,omitempty"`
}

type CosmosdbCassandraKeyspaceSpec struct {
	AccountName       string `json:"account_name"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
}

type CosmosdbCassandraKeyspaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CosmosdbCassandraKeyspaceList is a list of CosmosdbCassandraKeyspaces
type CosmosdbCassandraKeyspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CosmosdbCassandraKeyspace CRD objects
	Items []CosmosdbCassandraKeyspace `json:"items,omitempty"`
}
