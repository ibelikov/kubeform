package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermPostgresqlConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermPostgresqlConfigurationSpec   `json:"spec,omitempty"`
	Status            AzurermPostgresqlConfigurationStatus `json:"status,omitempty"`
}

type AzurermPostgresqlConfigurationSpec struct {
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	ServerName        string `json:"server_name"`
	Value             string `json:"value"`
}

type AzurermPostgresqlConfigurationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermPostgresqlConfigurationList is a list of AzurermPostgresqlConfigurations
type AzurermPostgresqlConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermPostgresqlConfiguration CRD objects
	Items []AzurermPostgresqlConfiguration `json:"items,omitempty"`
}
