package verifier

import (
    "errors"
    "encoding/json"
    "github.com/kcbojanowski/aws-iam-policy-verifier/pkg/model"
)

func VerifyPolicyJson(policyJson []byte) (bool, error) {
    var policy model.IAMPolicy
    if err := json.Unmarshal(policyJson, &policy); err != nil {
        return false, err
    }

    // Validate format using the ValidateFormat
    if err := ValidateFormat(policy); err != nil {
        return false, err
    }

    // Check for asterisk in Resource fields
    if containsAsterisk(policy) {
        return false, errors.New("use of '*' in Resource is not allowed")
    }

    return true, nil
}

func containsAsterisk(policy model.IAMPolicy) bool {
    for _, stmt := range policy.PolicyDocument.Statements {
        for _, res := range stmt.Resource {
            if res == "*" {
                return true
            }
        }
    }
    return false
}
