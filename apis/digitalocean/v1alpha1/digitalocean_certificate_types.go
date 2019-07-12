package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DigitaloceanCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanCertificateSpec   `json:"spec,omitempty"`
	Status            DigitaloceanCertificateStatus `json:"status,omitempty"`
}

type DigitaloceanCertificateSpec struct {
	Name             string   `json:"name"`
	CertificateChain string   `json:"certificate_chain"`
	Domains          []string `json:"domains"`
	Type             string   `json:"type"`
	State            string   `json:"state"`
	NotAfter         string   `json:"not_after"`
	PrivateKey       string   `json:"private_key"`
	LeafCertificate  string   `json:"leaf_certificate"`
	Sha1Fingerprint  string   `json:"sha1_fingerprint"`
}

type DigitaloceanCertificateStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DigitaloceanCertificateList is a list of DigitaloceanCertificates
type DigitaloceanCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanCertificate CRD objects
	Items []DigitaloceanCertificate `json:"items,omitempty"`
}
