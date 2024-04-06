package verifier

import (
    "encoding/json"
)

// structure of the AWS IAM Policy 
type PolicyDocument struct {
    Version   string      `json:"Version"`
    Statement []Statement `json:"Statement"`
}

type Statement struct {
    Sid      string   `json:"Sid"`
    Effect   string   `json:"Effect"`
    Action   []string `json:"Action"`
    Resource string   `json:"Resource"`
}

type IAMPolicy struct {
    PolicyName     string         `json:"PolicyName"`
    PolicyDocument PolicyDocument `json:"PolicyDocument"`
}


func VerifyPolicyJSON(policyJSON []byte) (bool, error) {
    var policy IAMPolicy

    err := json.Unmarshal(policyJSON, &policy)
    if err != nil {
        return false, err
    }

    for _, statement := range policy.PolicyDocument.Statement {
        if statement.Resource == "*" {
            return false, nil
        }
    }

    return true, nil
}
