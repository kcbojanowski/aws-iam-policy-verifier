package verifier

import (
    "github.com/kcbojanowski/aws-iam-policy-verifier/pkg/verifier/format"
    "github.com/kcbojanowski/aws-iam-policy-verifier/pkg/verifier/resource"
)

func ValidatePolicyJson(policyJson []byte) (bool, error) {
    policy, err := format.ValidateFormat(policyJson)
    if err != nil {
        return false, err
    }

    validResource, err := resource.ValidateResource(policy)
    if err != nil {
        return false, err
    }

    return validResource, nil
}
