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

type SecurityhubStandardsSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityhubStandardsSubscriptionSpec   `json:"spec,omitempty"`
	Status            SecurityhubStandardsSubscriptionStatus `json:"status,omitempty"`
}

type SecurityhubStandardsSubscriptionSpec struct {
	StandardsArn string `json:"standards_arn"`
}

type SecurityhubStandardsSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SecurityhubStandardsSubscriptionList is a list of SecurityhubStandardsSubscriptions
type SecurityhubStandardsSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SecurityhubStandardsSubscription CRD objects
	Items []SecurityhubStandardsSubscription `json:"items,omitempty"`
}
