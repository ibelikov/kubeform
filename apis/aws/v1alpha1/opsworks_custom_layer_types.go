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

type OpsworksCustomLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksCustomLayerSpec   `json:"spec,omitempty"`
	Status            OpsworksCustomLayerStatus `json:"status,omitempty"`
}

type OpsworksCustomLayerSpecEbsVolume struct {
	// +optional
	Iops          int    `json:"iops,omitempty"`
	MountPoint    string `json:"mount_point"`
	NumberOfDisks int    `json:"number_of_disks"`
	// +optional
	RaidLevel string `json:"raid_level,omitempty"`
	Size      int    `json:"size"`
	// +optional
	Type string `json:"type,omitempty"`
}

type OpsworksCustomLayerSpec struct {
	// +optional
	AutoAssignElasticIps bool `json:"auto_assign_elastic_ips,omitempty"`
	// +optional
	AutoAssignPublicIps bool `json:"auto_assign_public_ips,omitempty"`
	// +optional
	AutoHealing bool `json:"auto_healing,omitempty"`
	// +optional
	CustomConfigureRecipes []string `json:"custom_configure_recipes,omitempty"`
	// +optional
	CustomDeployRecipes []string `json:"custom_deploy_recipes,omitempty"`
	// +optional
	CustomInstanceProfileArn string `json:"custom_instance_profile_arn,omitempty"`
	// +optional
	CustomJson string `json:"custom_json,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	CustomSecurityGroupIds []string `json:"custom_security_group_ids,omitempty"`
	// +optional
	CustomSetupRecipes []string `json:"custom_setup_recipes,omitempty"`
	// +optional
	CustomShutdownRecipes []string `json:"custom_shutdown_recipes,omitempty"`
	// +optional
	CustomUndeployRecipes []string `json:"custom_undeploy_recipes,omitempty"`
	// +optional
	DrainElbOnShutdown bool `json:"drain_elb_on_shutdown,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EbsVolume *[]OpsworksCustomLayerSpec `json:"ebs_volume,omitempty"`
	// +optional
	ElasticLoadBalancer string `json:"elastic_load_balancer,omitempty"`
	// +optional
	InstallUpdatesOnBoot bool `json:"install_updates_on_boot,omitempty"`
	// +optional
	InstanceShutdownTimeout int    `json:"instance_shutdown_timeout,omitempty"`
	Name                    string `json:"name"`
	ShortName               string `json:"short_name"`
	StackId                 string `json:"stack_id"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SystemPackages []string `json:"system_packages,omitempty"`
	// +optional
	UseEbsOptimizedInstances bool `json:"use_ebs_optimized_instances,omitempty"`
}

type OpsworksCustomLayerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OpsworksCustomLayerList is a list of OpsworksCustomLayers
type OpsworksCustomLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpsworksCustomLayer CRD objects
	Items []OpsworksCustomLayer `json:"items,omitempty"`
}
