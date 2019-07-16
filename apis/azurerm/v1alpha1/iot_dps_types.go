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

type IotDps struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IotDpsSpec   `json:"spec,omitempty"`
	Status            IotDpsStatus `json:"status,omitempty"`
}

type IotDpsSpecSku struct {
	Capacity int    `json:"capacity"`
	Name     string `json:"name"`
	Tier     string `json:"tier"`
}

type IotDpsSpec struct {
	Location          string `json:"location"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Sku []IotDpsSpec `json:"sku"`
}

type IotDpsStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IotDpsList is a list of IotDpss
type IotDpsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IotDps CRD objects
	Items []IotDps `json:"items,omitempty"`
}
