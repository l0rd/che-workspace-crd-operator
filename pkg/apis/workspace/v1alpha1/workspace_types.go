package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Workspace is the Schema for the workspaces API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name=Id,type=string,JSONPath=.status.workspaceId
// +kubebuilder:printcolumn:name=Enabled,type=boolean,JSONPath=.spec.started
// +kubebuilder:printcolumn:name=Status,type=string,JSONPath=.status.phase
// +kubebuilder:printcolumn:name=Url,type=string,JSONPath=.status.ideUrl
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Desired state of the workspace
	Spec   WorkspaceSpec   `json:"spec,omitempty"`
	// Observed state of the workspace
	Status WorkspaceStatus `json:"status,omitempty"`
}

// WorkspaceSpec defines the desired state of Workspace
// +k8s:openapi-gen=true
type WorkspaceSpec struct {
	// Whether the workspace should be started or stopped
	Started bool          `json:"started"`
	// Exposure class the defines how the workspace will be exposed toon the external network
	ExposureClass string  `json:"exposureClass,omitempty"`
	// Workspace Structure defined in the Devfile format syntax.
	// For more details see the Che 7 documentation: https://www.eclipse.org/che/docs/che-7/making-a-workspace-portable-using-a-devfile/
	Devfile DevFileSpec   `json:"devfile"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkspaceList contains a list of Workspace
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
