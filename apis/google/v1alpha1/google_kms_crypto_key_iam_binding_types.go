package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleKmsCryptoKeyIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleKmsCryptoKeyIamBindingSpec   `json:"spec,omitempty"`
	Status            GoogleKmsCryptoKeyIamBindingStatus `json:"status,omitempty"`
}

type GoogleKmsCryptoKeyIamBindingSpec struct {
	CryptoKeyId string   `json:"crypto_key_id"`
	Etag        string   `json:"etag"`
	Role        string   `json:"role"`
	Members     []string `json:"members"`
}

type GoogleKmsCryptoKeyIamBindingStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleKmsCryptoKeyIamBindingList is a list of GoogleKmsCryptoKeyIamBindings
type GoogleKmsCryptoKeyIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleKmsCryptoKeyIamBinding CRD objects
	Items []GoogleKmsCryptoKeyIamBinding `json:"items,omitempty"`
}
