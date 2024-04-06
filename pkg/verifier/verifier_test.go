package verifier

import (
    "testing"
)

func TestVerifyPolicyJSON(t *testing.T) {
    testPolicyJSON := `{
        "PolicyName": "root",
        "PolicyDocument": {
            "Version": "2012-10-17",
            "Statement": [
                {
                    "Sid": "TestStatement",
                    "Effect": "Allow",
                    "Action": [
                        "iam:ListRoles",
                        "iam:ListUsers"
                    ],
                    "Resource": "*"
                }
            ]
        }
    }`

    policyBytes := []byte(testPolicyJSON)

    valid, err := VerifyPolicyJSON(policyBytes)
    if err != nil {
        t.Errorf("VerifyPolicyJSON returned an unexpected error: %v", err)
    }
    if valid {
        t.Errorf("VerifyPolicyJSON should have returned false for a policy with a Resource of '*'")
    }
}
