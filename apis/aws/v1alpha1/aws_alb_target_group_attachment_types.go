package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAlbTargetGroupAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAlbTargetGroupAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsAlbTargetGroupAttachmentStatus `json:"status,omitempty"`
}

type AwsAlbTargetGroupAttachmentSpec struct {
	TargetGroupArn   string `json:"target_group_arn"`
	TargetId         string `json:"target_id"`
	Port             int    `json:"port"`
	AvailabilityZone string `json:"availability_zone"`
}

type AwsAlbTargetGroupAttachmentStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbTargetGroupAttachmentList is a list of AwsAlbTargetGroupAttachments
type AwsAlbTargetGroupAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAlbTargetGroupAttachment CRD objects
	Items []AwsAlbTargetGroupAttachment `json:"items,omitempty"`
}
