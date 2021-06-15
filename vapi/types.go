package vapi

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type AdmissionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AdmissionPolicySpec `json:"spec"`
}

type AdmissionPolicySpec struct {
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
