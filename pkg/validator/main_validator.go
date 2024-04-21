package validator

import (
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
		fmt.Println("Decoder: Invalid format of JSON ❌")
		return false, err
	}

	err = ValidateIAMPolicy(policy)
	if err != nil {
		fmt.Println("Validator: Invalid IAM policy format ❌")
		return false, err
	}

	fmt.Println("Policy JSON validation passed ✅")
	return true, nil
}
