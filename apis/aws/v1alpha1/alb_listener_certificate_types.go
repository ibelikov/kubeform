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

type AlbListenerCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlbListenerCertificateSpec   `json:"spec,omitempty"`
	Status            AlbListenerCertificateStatus `json:"status,omitempty"`
}

type AlbListenerCertificateSpec struct {
	CertificateArn string `json:"certificate_arn"`
	ListenerArn    string `json:"listener_arn"`
}

type AlbListenerCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AlbListenerCertificateList is a list of AlbListenerCertificates
type AlbListenerCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AlbListenerCertificate CRD objects
	Items []AlbListenerCertificate `json:"items,omitempty"`
}
