package vapi

var AdmissionPolicies []AdmissionPolicy

func Startup() {
	LoadPolicyFromCustomResources()

}
