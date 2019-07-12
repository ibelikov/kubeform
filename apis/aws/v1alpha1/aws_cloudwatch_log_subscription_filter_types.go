package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudwatchLogSubscriptionFilter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudwatchLogSubscriptionFilterSpec   `json:"spec,omitempty"`
	Status            AwsCloudwatchLogSubscriptionFilterStatus `json:"status,omitempty"`
}

type AwsCloudwatchLogSubscriptionFilterSpec struct {
	Name           string `json:"name"`
	DestinationArn string `json:"destination_arn"`
	FilterPattern  string `json:"filter_pattern"`
	LogGroupName   string `json:"log_group_name"`
	RoleArn        string `json:"role_arn"`
	Distribution   string `json:"distribution"`
}

type AwsCloudwatchLogSubscriptionFilterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchLogSubscriptionFilterList is a list of AwsCloudwatchLogSubscriptionFilters
type AwsCloudwatchLogSubscriptionFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudwatchLogSubscriptionFilter CRD objects
	Items []AwsCloudwatchLogSubscriptionFilter `json:"items,omitempty"`
}
