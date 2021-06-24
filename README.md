# jpe

## Problem statement

jpe is a Kubernetes native policy engine to validate creation and update of objects into the kubernetes cluster. It has been developed in response to difficult to learn policy engines like OPA and kyverno.

You write your policies in pure javascript, a language that is easy to understand, something you probably already know and there is tons of documentation available.

You do not have to learn a separate language like rego (in OPA) or try to put logic in YAML (in kyverno)

## How it works

A sample policy looks like.

```
apiVersion: jpe.prabhatsharma.in/v1alpha1
kind: AdmissionPolicy
metadata:
  name: simplepolicy
spec:
  rules:
  - name: rule1
    resourceKind: pod
    validationFailureAction: enforce
    description: Any pods that have the name simple are really weird and we should prevent them from running. 
    message: Pod name should not be simple.
    rule: |
      function validate(resource){
        resource = JSON.parse(resource);
        if(resource.metadata.name != "simple") {
          return true;
        };
        return false;
      }

```

You will get the object that you want to validate as a JSON object. You can write the validate function to validate any part of it, any way you choose to do so.


# Installation


