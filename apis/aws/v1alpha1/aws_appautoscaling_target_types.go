package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppautoscalingTarget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppautoscalingTargetSpec   `json:"spec,omitempty"`
	Status            AwsAppautoscalingTargetStatus `json:"status,omitempty"`
}

type AwsAppautoscalingTargetSpec struct {
	ScalableDimension string `json:"scalable_dimension"`
	ServiceNamespace  string `json:"service_namespace"`
	MaxCapacity       int    `json:"max_capacity"`
	MinCapacity       int    `json:"min_capacity"`
	ResourceId        string `json:"resource_id"`
	RoleArn           string `json:"role_arn"`
}

type AwsAppautoscalingTargetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppautoscalingTargetList is a list of AwsAppautoscalingTargets
type AwsAppautoscalingTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppautoscalingTarget CRD objects
	Items []AwsAppautoscalingTarget `json:"items,omitempty"`
}
