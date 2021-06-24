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
	var rr RuleResponse
	c.BindJSON(&aReviewRequest)
	if strings.ToLower(aReviewRequest.Request.Kind.Kind) == "admissionpolicy" {
		// This is an admission policy
		if aReviewRequest.Request.Operation == "DELETE" {
			// Delete the policy from memory
			rr = DeletePolicy(&aReviewRequest)
		} else if aReviewRequest.Request.Operation == "CREATE" { // CREATE
			// Load this policy in memory to use along with existing policies.
			rr = AddToExistingPolicies(&aReviewRequest)
		} else if aReviewRequest.Request.Operation == "UPDATE" { // CREATE
			// Update policy in memory to use along with existing policies.
			rr = UpdatePolicy(&aReviewRequest)
		}

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
		if aReviewRequest.Request.Operation == "DELETE" {
			deleteResponse := &v1.AdmissionReview{
				TypeMeta: aReviewRequest.TypeMeta,
				Response: &v1.AdmissionResponse{
					UID:     aReviewRequest.Request.UID,
					Allowed: true,
					Result: &metav1.Status{
						Status:  "Success",
						Message: "Skipping DELETE",
					},
				},
			}

			c.JSON(200, deleteResponse)

			return

		}
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

		finalResponse := v1.AdmissionReview{
			TypeMeta: aReviewRequest.TypeMeta,
			Response: &aReviewResponse,
		}

		if !rr.Allowed {
			aRR, _ := json.Marshal(aReviewRequest)
			arV, _ := json.Marshal(finalResponse)
			fmt.Println("ReviewRequest is: ", string(aRR))
			fmt.Println("ReviewResponse is: ", string(arV))
		}

		c.JSON(200, finalResponse)
	}
}
