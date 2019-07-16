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

type DnsManagedZone struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsManagedZoneSpec   `json:"spec,omitempty"`
	Status            DnsManagedZoneStatus `json:"status,omitempty"`
}

type DnsManagedZoneSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	DnsName     string `json:"dns_name"`
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	Name   string            `json:"name"`
}

type DnsManagedZoneStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DnsManagedZoneList is a list of DnsManagedZones
type DnsManagedZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DnsManagedZone CRD objects
	Items []DnsManagedZone `json:"items,omitempty"`
}
