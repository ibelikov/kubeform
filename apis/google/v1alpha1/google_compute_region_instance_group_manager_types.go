package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeRegionInstanceGroupManager struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeRegionInstanceGroupManagerSpec   `json:"spec,omitempty"`
	Status            GoogleComputeRegionInstanceGroupManagerStatus `json:"status,omitempty"`
}

type GoogleComputeRegionInstanceGroupManagerSpecRollingUpdatePolicy struct {
	MinimalAction         string `json:"minimal_action"`
	Type                  string `json:"type"`
	MaxSurgeFixed         int    `json:"max_surge_fixed"`
	MaxSurgePercent       int    `json:"max_surge_percent"`
	MaxUnavailableFixed   int    `json:"max_unavailable_fixed"`
	MaxUnavailablePercent int    `json:"max_unavailable_percent"`
	MinReadySec           int    `json:"min_ready_sec"`
}

type GoogleComputeRegionInstanceGroupManagerSpecAutoHealingPolicies struct {
	HealthCheck     string `json:"health_check"`
	InitialDelaySec int    `json:"initial_delay_sec"`
}

type GoogleComputeRegionInstanceGroupManagerSpecVersionTargetSize struct {
	Fixed   int `json:"fixed"`
	Percent int `json:"percent"`
}

type GoogleComputeRegionInstanceGroupManagerSpecVersion struct {
	Name             string                                               `json:"name"`
	InstanceTemplate string                                               `json:"instance_template"`
	TargetSize       []GoogleComputeRegionInstanceGroupManagerSpecVersion `json:"target_size"`
}

type GoogleComputeRegionInstanceGroupManagerSpecNamedPort struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

type GoogleComputeRegionInstanceGroupManagerSpec struct {
	Description             string                                        `json:"description"`
	RollingUpdatePolicy     []GoogleComputeRegionInstanceGroupManagerSpec `json:"rolling_update_policy"`
	InstanceTemplate        string                                        `json:"instance_template"`
	TargetPools             []string                                      `json:"target_pools"`
	AutoHealingPolicies     []GoogleComputeRegionInstanceGroupManagerSpec `json:"auto_healing_policies"`
	BaseInstanceName        string                                        `json:"base_instance_name"`
	Fingerprint             string                                        `json:"fingerprint"`
	InstanceGroup           string                                        `json:"instance_group"`
	SelfLink                string                                        `json:"self_link"`
	Project                 string                                        `json:"project"`
	UpdateStrategy          string                                        `json:"update_strategy"`
	TargetSize              int                                           `json:"target_size"`
	WaitForInstances        bool                                          `json:"wait_for_instances"`
	Version                 []GoogleComputeRegionInstanceGroupManagerSpec `json:"version"`
	Name                    string                                        `json:"name"`
	Region                  string                                        `json:"region"`
	NamedPort               []GoogleComputeRegionInstanceGroupManagerSpec `json:"named_port"`
	DistributionPolicyZones []string                                      `json:"distribution_policy_zones"`
}

type GoogleComputeRegionInstanceGroupManagerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeRegionInstanceGroupManagerList is a list of GoogleComputeRegionInstanceGroupManagers
type GoogleComputeRegionInstanceGroupManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeRegionInstanceGroupManager CRD objects
	Items []GoogleComputeRegionInstanceGroupManager `json:"items,omitempty"`
}
