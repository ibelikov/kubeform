package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleProjectOrganizationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleProjectOrganizationPolicySpec   `json:"spec,omitempty"`
	Status            GoogleProjectOrganizationPolicyStatus `json:"status,omitempty"`
}

type GoogleProjectOrganizationPolicySpecRestorePolicy struct {
	Default bool `json:"default"`
}

type GoogleProjectOrganizationPolicySpecBooleanPolicy struct {
	Enforced bool `json:"enforced"`
}

type GoogleProjectOrganizationPolicySpecListPolicyAllow struct {
	All    bool     `json:"all"`
	Values []string `json:"values"`
}

type GoogleProjectOrganizationPolicySpecListPolicyDeny struct {
	All    bool     `json:"all"`
	Values []string `json:"values"`
}

type GoogleProjectOrganizationPolicySpecListPolicy struct {
	Allow          []GoogleProjectOrganizationPolicySpecListPolicy `json:"allow"`
	Deny           []GoogleProjectOrganizationPolicySpecListPolicy `json:"deny"`
	SuggestedValue string                                          `json:"suggested_value"`
}

type GoogleProjectOrganizationPolicySpec struct {
	Version       int                                   `json:"version"`
	Etag          string                                `json:"etag"`
	UpdateTime    string                                `json:"update_time"`
	RestorePolicy []GoogleProjectOrganizationPolicySpec `json:"restore_policy"`
	Project       string                                `json:"project"`
	Constraint    string                                `json:"constraint"`
	BooleanPolicy []GoogleProjectOrganizationPolicySpec `json:"boolean_policy"`
	ListPolicy    []GoogleProjectOrganizationPolicySpec `json:"list_policy"`
}

type GoogleProjectOrganizationPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleProjectOrganizationPolicyList is a list of GoogleProjectOrganizationPolicys
type GoogleProjectOrganizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleProjectOrganizationPolicy CRD objects
	Items []GoogleProjectOrganizationPolicy `json:"items,omitempty"`
}
