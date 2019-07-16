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

type MediaPackageChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MediaPackageChannelSpec   `json:"spec,omitempty"`
	Status            MediaPackageChannelStatus `json:"status,omitempty"`
}

type MediaPackageChannelSpec struct {
	ChannelId string `json:"channel_id"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type MediaPackageChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MediaPackageChannelList is a list of MediaPackageChannels
type MediaPackageChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MediaPackageChannel CRD objects
	Items []MediaPackageChannel `json:"items,omitempty"`
}