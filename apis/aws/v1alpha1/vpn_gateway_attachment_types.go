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

type VpnGatewayAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpnGatewayAttachmentSpec   `json:"spec,omitempty"`
	Status            VpnGatewayAttachmentStatus `json:"status,omitempty"`
}

type VpnGatewayAttachmentSpec struct {
	VpcId        string `json:"vpc_id"`
	VpnGatewayId string `json:"vpn_gateway_id"`
}

type VpnGatewayAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpnGatewayAttachmentList is a list of VpnGatewayAttachments
type VpnGatewayAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpnGatewayAttachment CRD objects
	Items []VpnGatewayAttachment `json:"items,omitempty"`
}
