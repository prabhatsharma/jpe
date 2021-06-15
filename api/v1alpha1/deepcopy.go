package v1alpha1

import "k8s.io/apimachinery/pkg/runtime"

// DeepCopyInto copies all properties of this object into another object of the
// same type that is provided as a pointer.
func (in *AdmissionPolicy) DeepCopyInto(out *AdmissionPolicy) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = AdmissionPolicySpec{
		Rules: in.Spec.Rules,
	}
}

// DeepCopyObject returns a generically typed copy of an object
func (in *AdmissionPolicy) DeepCopyObject() runtime.Object {
	out := AdmissionPolicy{}
	in.DeepCopyInto(&out)

	return &out
}

// DeepCopyObject returns a generically typed copy of an object
func (in *AdmissionPolicyList) DeepCopyObject() runtime.Object {
	out := AdmissionPolicyList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]AdmissionPolicy, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
