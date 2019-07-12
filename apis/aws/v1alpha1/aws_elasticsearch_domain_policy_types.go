package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticsearchDomainPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElasticsearchDomainPolicySpec   `json:"spec,omitempty"`
	Status            AwsElasticsearchDomainPolicyStatus `json:"status,omitempty"`
}

type AwsElasticsearchDomainPolicySpec struct {
	DomainName     string `json:"domain_name"`
	AccessPolicies string `json:"access_policies"`
}

type AwsElasticsearchDomainPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticsearchDomainPolicyList is a list of AwsElasticsearchDomainPolicys
type AwsElasticsearchDomainPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElasticsearchDomainPolicy CRD objects
	Items []AwsElasticsearchDomainPolicy `json:"items,omitempty"`
}
