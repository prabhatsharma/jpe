package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//go:generate controller-gen object paths=$GOFILE

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AdmissionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AdmissionPolicySpec `json:"spec"`
}

type AdmissionPolicySpec struct {
	// Rules is a collection of one or more independent rules against which resources will be validated against.
	Rules []Rule `json:"rules"`
}

type Rule struct {
	Name                    string `json:"name"`
	Resource                string `json:"resource"`
	ValidationFailureAction string `json:"validationFailureAction"`
	Rule                    string `json:"rule"`
	Description             string `json:"description"`
	Message                 string `json:"message"`
}

type RuleResponse struct {
	Allowed bool   `json:"allowed"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AdmissionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AdmissionPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AdmissionPolicy{}, &AdmissionPolicyList{})
}
