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

type OrganizationIamCustomRole struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationIamCustomRoleSpec   `json:"spec,omitempty"`
	Status            OrganizationIamCustomRoleStatus `json:"status,omitempty"`
}

type OrganizationIamCustomRoleSpec struct {
	// +optional
	// Deprecated
	Deleted bool `json:"deleted,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	OrgId       string `json:"org_id"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Permissions []string `json:"permissions"`
	RoleId      string   `json:"role_id"`
	// +optional
	Stage string `json:"stage,omitempty"`
	Title string `json:"title"`
}

type OrganizationIamCustomRoleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OrganizationIamCustomRoleList is a list of OrganizationIamCustomRoles
type OrganizationIamCustomRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OrganizationIamCustomRole CRD objects
	Items []OrganizationIamCustomRole `json:"items,omitempty"`
}
