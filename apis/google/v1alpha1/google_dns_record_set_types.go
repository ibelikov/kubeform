package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleDnsRecordSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleDnsRecordSetSpec   `json:"spec,omitempty"`
	Status            GoogleDnsRecordSetStatus `json:"status,omitempty"`
}

type GoogleDnsRecordSetSpec struct {
	ManagedZone string   `json:"managed_zone"`
	Name        string   `json:"name"`
	Rrdatas     []string `json:"rrdatas"`
	Ttl         int      `json:"ttl"`
	Type        string   `json:"type"`
	Project     string   `json:"project"`
}

type GoogleDnsRecordSetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleDnsRecordSetList is a list of GoogleDnsRecordSets
type GoogleDnsRecordSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleDnsRecordSet CRD objects
	Items []GoogleDnsRecordSet `json:"items,omitempty"`
}
