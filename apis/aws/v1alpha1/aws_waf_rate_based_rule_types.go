package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsWafRateBasedRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWafRateBasedRuleSpec   `json:"spec,omitempty"`
	Status            AwsWafRateBasedRuleStatus `json:"status,omitempty"`
}

type AwsWafRateBasedRuleSpecPredicates struct {
	DataId  string `json:"data_id"`
	Type    string `json:"type"`
	Negated bool   `json:"negated"`
}

type AwsWafRateBasedRuleSpec struct {
	Predicates []AwsWafRateBasedRuleSpec `json:"predicates"`
	RateKey    string                    `json:"rate_key"`
	RateLimit  int                       `json:"rate_limit"`
	Name       string                    `json:"name"`
	MetricName string                    `json:"metric_name"`
}

type AwsWafRateBasedRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsWafRateBasedRuleList is a list of AwsWafRateBasedRules
type AwsWafRateBasedRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWafRateBasedRule CRD objects
	Items []AwsWafRateBasedRule `json:"items,omitempty"`
}
