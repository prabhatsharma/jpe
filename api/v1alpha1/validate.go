package v1alpha1

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Validate(c *gin.Context) {
	var aReviewRequest v1.AdmissionReview
	c.BindJSON(&aReviewRequest)
	if strings.ToLower(aReviewRequest.Request.Kind.Kind) == "admissionpolicy" {
		// This is an admission policy. Load this policy in memory to use along with existing policies.
		rr := AddToExistingPolicies(&aReviewRequest)

		status := &metav1.Status{
			Status:  rr.Status, // Success or Failure
			Message: aReviewRequest.Request.Name + " : " + rr.Message,
		}

		// This is the response that is returned to k8s
		aReviewResponse := v1.AdmissionResponse{
			UID:     aReviewRequest.Request.UID,
			Allowed: true, // Set this to true for when validation succeeds. Set this to false for failed validation
			Result:  status,
		}

		c.JSON(200, aReviewResponse)

	} else {
		// Validate the resource based on existing admission policies
		rr := ValidateResource(&aReviewRequest)

		status := &metav1.Status{
			Status:  rr.Status, // Success or Failure
			Message: rr.Message,
		}

		// This is the response that is returned to k8s
		aReviewResponse := v1.AdmissionResponse{
			UID:     aReviewRequest.Request.UID,
			Allowed: rr.Allowed, // Set this to true for when validation succeeds. Set this to false for failed validation
			Result:  status,
		}

		if !rr.Allowed {
			aRR, _ := json.Marshal(aReviewRequest)
			arV, _ := json.Marshal(aReviewResponse)
			fmt.Println("ReviewRequest is: ", string(aRR))
			fmt.Println("ReviewResponse is: ", string(arV))
		}

		finalResponse := v1.AdmissionReview{
			TypeMeta: aReviewRequest.TypeMeta,
			Response: &aReviewResponse,
		}

		c.JSON(200, finalResponse)
	}
}
