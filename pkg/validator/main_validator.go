package validator

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func ValidatePolicyJson(path string) (bool, error) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading %s: %s ❌\n", path, err)
		return false, err
	}

	policy, err := loadPolicyFromJSON(fileContent)
	if err != nil {
		fmt.Println("Invalid format of JSON ❌")
		return false, err
	}

	err = ValidateIAMPolicy(policy)
	if err != nil {
		fmt.Println("Invalid IAM policy format ❌")
		return false, err
	}

	for _, statement := range policy.PolicyDocument.Statements {
		for _, resource := range statement.Resource {
			if resource == "*" {
				return false, errors.New("resource field contains a single asterisk")
			}
		}
	}

	return true, nil
}
