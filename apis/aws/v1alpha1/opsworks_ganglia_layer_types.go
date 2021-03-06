/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type OpsworksGangliaLayer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OpsworksGangliaLayerSpec   `json:"spec,omitempty"`
	Status            OpsworksGangliaLayerStatus `json:"status,omitempty"`
}

type OpsworksGangliaLayerSpecEbsVolume struct {
	// +optional
	Iops          int64  `json:"iops,omitempty" tf:"iops,omitempty"`
	MountPoint    string `json:"mountPoint" tf:"mount_point"`
	NumberOfDisks int64  `json:"numberOfDisks" tf:"number_of_disks"`
	// +optional
	RaidLevel string `json:"raidLevel,omitempty" tf:"raid_level,omitempty"`
	Size      int64  `json:"size" tf:"size"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type OpsworksGangliaLayerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutoAssignElasticIPS bool `json:"autoAssignElasticIPS,omitempty" tf:"auto_assign_elastic_ips,omitempty"`
	// +optional
	AutoAssignPublicIPS bool `json:"autoAssignPublicIPS,omitempty" tf:"auto_assign_public_ips,omitempty"`
	// +optional
	AutoHealing bool `json:"autoHealing,omitempty" tf:"auto_healing,omitempty"`
	// +optional
	CustomConfigureRecipes []string `json:"customConfigureRecipes,omitempty" tf:"custom_configure_recipes,omitempty"`
	// +optional
	CustomDeployRecipes []string `json:"customDeployRecipes,omitempty" tf:"custom_deploy_recipes,omitempty"`
	// +optional
	CustomInstanceProfileArn string `json:"customInstanceProfileArn,omitempty" tf:"custom_instance_profile_arn,omitempty"`
	// +optional
	CustomJSON string `json:"customJSON,omitempty" tf:"custom_json,omitempty"`
	// +optional
	CustomSecurityGroupIDS []string `json:"customSecurityGroupIDS,omitempty" tf:"custom_security_group_ids,omitempty"`
	// +optional
	CustomSetupRecipes []string `json:"customSetupRecipes,omitempty" tf:"custom_setup_recipes,omitempty"`
	// +optional
	CustomShutdownRecipes []string `json:"customShutdownRecipes,omitempty" tf:"custom_shutdown_recipes,omitempty"`
	// +optional
	CustomUndeployRecipes []string `json:"customUndeployRecipes,omitempty" tf:"custom_undeploy_recipes,omitempty"`
	// +optional
	DrainElbOnShutdown bool `json:"drainElbOnShutdown,omitempty" tf:"drain_elb_on_shutdown,omitempty"`
	// +optional
	EbsVolume []OpsworksGangliaLayerSpecEbsVolume `json:"ebsVolume,omitempty" tf:"ebs_volume,omitempty"`
	// +optional
	ElasticLoadBalancer string `json:"elasticLoadBalancer,omitempty" tf:"elastic_load_balancer,omitempty"`
	// +optional
	InstallUpdatesOnBoot bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot,omitempty"`
	// +optional
	InstanceShutdownTimeout int64 `json:"instanceShutdownTimeout,omitempty" tf:"instance_shutdown_timeout,omitempty"`
	// +optional
	Name     string `json:"name,omitempty" tf:"name,omitempty"`
	Password string `json:"password" tf:"password"`
	StackID  string `json:"stackID" tf:"stack_id"`
	// +optional
	SystemPackages []string `json:"systemPackages,omitempty" tf:"system_packages,omitempty"`
	// +optional
	Url string `json:"url,omitempty" tf:"url,omitempty"`
	// +optional
	UseEbsOptimizedInstances bool `json:"useEbsOptimizedInstances,omitempty" tf:"use_ebs_optimized_instances,omitempty"`
	// +optional
	Username string `json:"username,omitempty" tf:"username,omitempty"`
}

type OpsworksGangliaLayerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *OpsworksGangliaLayerSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OpsworksGangliaLayerList is a list of OpsworksGangliaLayers
type OpsworksGangliaLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OpsworksGangliaLayer CRD objects
	Items []OpsworksGangliaLayer `json:"items,omitempty"`
}
