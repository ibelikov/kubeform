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

type NodebalancerConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodebalancerConfigSpec   `json:"spec,omitempty"`
	Status            NodebalancerConfigStatus `json:"status,omitempty"`
}

type NodebalancerConfigSpecNodeStatus struct {
	StatusUp   int `json:"status_up"`
	StatusDown int `json:"status_down"`
}

type NodebalancerConfigSpec struct {
	CheckAttempts  int               `json:"check_attempts"`
	Stickiness     string            `json:"stickiness"`
	CheckPath      string            `json:"check_path"`
	CipherSuite    string            `json:"cipher_suite"`
	SslCommonname  string            `json:"ssl_commonname"`
	NodeStatus     map[string]string `json:"node_status"`
	Protocol       string            `json:"protocol"`
	Algorithm      string            `json:"algorithm"`
	CheckBody      string            `json:"check_body"`
	CheckPassive   bool              `json:"check_passive"`
	SslKey         string            `json:"ssl_key"`
	NodebalancerId int               `json:"nodebalancer_id"`
	Port           int               `json:"port"`
	CheckTimeout   int               `json:"check_timeout"`
	CheckInterval  int               `json:"check_interval"`
	Check          string            `json:"check"`
	SslFingerprint string            `json:"ssl_fingerprint"`
	SslCert        string            `json:"ssl_cert"`
}



type NodebalancerConfigStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NodebalancerConfigList is a list of NodebalancerConfigs
type NodebalancerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NodebalancerConfig CRD objects
	Items []NodebalancerConfig `json:"items,omitempty"`
}