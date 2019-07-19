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

type ElasticacheReplicationGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticacheReplicationGroupSpec   `json:"spec,omitempty"`
	Status            ElasticacheReplicationGroupStatus `json:"status,omitempty"`
}

type ElasticacheReplicationGroupSpecClusterMode struct {
	NumNodeGroups        int `json:"numNodeGroups" tf:"num_node_groups"`
	ReplicasPerNodeGroup int `json:"replicasPerNodeGroup" tf:"replicas_per_node_group"`
}

type ElasticacheReplicationGroupSpec struct {
	// +optional
	ApplyImmediately bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`
	// +optional
	AtRestEncryptionEnabled bool `json:"atRestEncryptionEnabled,omitempty" tf:"at_rest_encryption_enabled,omitempty"`
	// +optional
	// Sensitive Data. Provide secret name which contains one value only
	AuthToken *core.LocalObjectReference `json:"authToken,omitempty" tf:"auth_token,omitempty"`
	// +optional
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`
	// +optional
	AutomaticFailoverEnabled bool `json:"automaticFailoverEnabled,omitempty" tf:"automatic_failover_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ClusterMode []ElasticacheReplicationGroupSpecClusterMode `json:"clusterMode,omitempty" tf:"cluster_mode,omitempty"`
	// +optional
	Engine string `json:"engine,omitempty" tf:"engine,omitempty"`
	// +optional
	EngineVersion string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`
	// +optional
	MaintenanceWindow string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`
	// +optional
	NodeType string `json:"nodeType,omitempty" tf:"node_type,omitempty"`
	// +optional
	NotificationTopicArn string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn,omitempty"`
	// +optional
	NumberCacheClusters int `json:"numberCacheClusters,omitempty" tf:"number_cache_clusters,omitempty"`
	// +optional
	ParameterGroupName string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`
	// +optional
	Port                        int    `json:"port,omitempty" tf:"port,omitempty"`
	ReplicationGroupDescription string `json:"replicationGroupDescription" tf:"replication_group_description"`
	ReplicationGroupID          string `json:"replicationGroupID" tf:"replication_group_id"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupNames []string `json:"securityGroupNames,omitempty" tf:"security_group_names,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SnapshotArns []string `json:"snapshotArns,omitempty" tf:"snapshot_arns,omitempty"`
	// +optional
	SnapshotName string `json:"snapshotName,omitempty" tf:"snapshot_name,omitempty"`
	// +optional
	SnapshotRetentionLimit int `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit,omitempty"`
	// +optional
	SnapshotWindow string `json:"snapshotWindow,omitempty" tf:"snapshot_window,omitempty"`
	// +optional
	SubnetGroupName string `json:"subnetGroupName,omitempty" tf:"subnet_group_name,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TransitEncryptionEnabled bool                      `json:"transitEncryptionEnabled,omitempty" tf:"transit_encryption_enabled,omitempty"`
	ProviderRef              core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ElasticacheReplicationGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElasticacheReplicationGroupList is a list of ElasticacheReplicationGroups
type ElasticacheReplicationGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ElasticacheReplicationGroup CRD objects
	Items []ElasticacheReplicationGroup `json:"items,omitempty"`
}
