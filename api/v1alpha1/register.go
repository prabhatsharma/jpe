package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func AddKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(GroupVersion,
		&AdmissionPolicy{},
		&AdmissionPolicyList{},
	)

	metav1.AddToGroupVersion(scheme, GroupVersion)
	return nil
}
