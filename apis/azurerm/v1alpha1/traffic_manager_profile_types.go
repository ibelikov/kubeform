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

type TrafficManagerProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrafficManagerProfileSpec   `json:"spec,omitempty"`
	Status            TrafficManagerProfileStatus `json:"status,omitempty"`
}

type TrafficManagerProfileSpecDnsConfig struct {
	RelativeName string `json:"relative_name"`
	Ttl          int    `json:"ttl"`
}

type TrafficManagerProfileSpecMonitorConfig struct {
	// +optional
	Path     string `json:"path,omitempty"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type TrafficManagerProfileSpec struct {
	// +kubebuilder:validation:UniqueItems=true
	DnsConfig []TrafficManagerProfileSpec `json:"dns_config"`
	// +kubebuilder:validation:UniqueItems=true
	MonitorConfig        []TrafficManagerProfileSpec `json:"monitor_config"`
	Name                 string                      `json:"name"`
	ResourceGroupName    string                      `json:"resource_group_name"`
	TrafficRoutingMethod string                      `json:"traffic_routing_method"`
}

type TrafficManagerProfileStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// TrafficManagerProfileList is a list of TrafficManagerProfiles
type TrafficManagerProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TrafficManagerProfile CRD objects
	Items []TrafficManagerProfile `json:"items,omitempty"`
}
