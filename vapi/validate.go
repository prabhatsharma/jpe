package vapi

import (
	"strings"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Validate(c *gin.Context) {

	var aReviewRequest v1.AdmissionReview
	c.BindJSON(&aReviewRequest)
	if strings.ToLower(aReviewRequest.Request.Kind.Kind) == "admissionpolicy" {
		// This is an admission policy. Load this policy in memory to use.
		LoadPolicyFromAdmissionReview(&aReviewRequest)
	} else {
		// Validate the resource based on existing admission policies
		ValidateResource(&aReviewRequest)
	}

	status := &metav1.Status{
		Status:  "accepted",
		Reason:  "Will accept anything",
		Message: "Enjoy your life",
	}

	// This is the response that is returned to k8s
	aReviewResponse := v1.AdmissionResponse{
		UID:     aReviewRequest.Request.UID,
		Allowed: true, // Set this to true for when validation succeeds. Set this to false for failed validation
		Result:  status,
	}

	c.JSON(200, aReviewResponse)
}
