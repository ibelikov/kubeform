package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type HdinsightHbaseCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HdinsightHbaseClusterSpec   `json:"spec,omitempty"`
	Status            HdinsightHbaseClusterStatus `json:"status,omitempty"`
}

type HdinsightHbaseClusterSpecComponentVersion struct {
	Hbase string `json:"hbase" tf:"hbase"`
}

type HdinsightHbaseClusterSpecGateway struct {
	Enabled bool `json:"enabled" tf:"enabled"`
	// Sensitive Data. Provide secret name which contains one value only
	Password *core.LocalObjectReference `json:"password" tf:"password"`
	Username string                     `json:"username" tf:"username"`
}

type HdinsightHbaseClusterSpecRolesHeadNode struct {
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	Password *core.LocalObjectReference `json:"password,omitempty" tf:"password,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	Username string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id,omitempty"`
	VmSize           string `json:"vmSize" tf:"vm_size"`
}

type HdinsightHbaseClusterSpecRolesWorkerNode struct {
	// +optional
	MinInstanceCount int `json:"minInstanceCount,omitempty" tf:"min_instance_count,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	Password *core.LocalObjectReference `json:"password,omitempty" tf:"password,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	SubnetID            string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	TargetInstanceCount int    `json:"targetInstanceCount" tf:"target_instance_count"`
	Username            string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id,omitempty"`
	VmSize           string `json:"vmSize" tf:"vm_size"`
}

type HdinsightHbaseClusterSpecRolesZookeeperNode struct {
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	Password *core.LocalObjectReference `json:"password,omitempty" tf:"password,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	Username string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id,omitempty"`
	VmSize           string `json:"vmSize" tf:"vm_size"`
}

type HdinsightHbaseClusterSpecRoles struct {
	// +kubebuilder:validation:MaxItems=1
	HeadNode []HdinsightHbaseClusterSpecRolesHeadNode `json:"headNode" tf:"head_node"`
	// +kubebuilder:validation:MaxItems=1
	WorkerNode []HdinsightHbaseClusterSpecRolesWorkerNode `json:"workerNode" tf:"worker_node"`
	// +kubebuilder:validation:MaxItems=1
	ZookeeperNode []HdinsightHbaseClusterSpecRolesZookeeperNode `json:"zookeeperNode" tf:"zookeeper_node"`
}

type HdinsightHbaseClusterSpecStorageAccount struct {
	IsDefault bool `json:"isDefault" tf:"is_default"`
	// Sensitive Data. Provide secret name which contains one value only
	StorageAccountKey  *core.LocalObjectReference `json:"storageAccountKey" tf:"storage_account_key"`
	StorageContainerID string                     `json:"storageContainerID" tf:"storage_container_id"`
}

type HdinsightHbaseClusterSpec struct {
	ClusterVersion string `json:"clusterVersion" tf:"cluster_version"`
	// +kubebuilder:validation:MaxItems=1
	ComponentVersion []HdinsightHbaseClusterSpecComponentVersion `json:"componentVersion" tf:"component_version"`
	// +kubebuilder:validation:MaxItems=1
	Gateway           []HdinsightHbaseClusterSpecGateway `json:"gateway" tf:"gateway"`
	Location          string                             `json:"location" tf:"location"`
	Name              string                             `json:"name" tf:"name"`
	ResourceGroupName string                             `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Roles          []HdinsightHbaseClusterSpecRoles          `json:"roles" tf:"roles"`
	StorageAccount []HdinsightHbaseClusterSpecStorageAccount `json:"storageAccount" tf:"storage_account"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	Tier        string                    `json:"tier" tf:"tier"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type HdinsightHbaseClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// HdinsightHbaseClusterList is a list of HdinsightHbaseClusters
type HdinsightHbaseClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of HdinsightHbaseCluster CRD objects
	Items []HdinsightHbaseCluster `json:"items,omitempty"`
}
