package validator

/*
This file provides a function to validate the format of an IAM policy.
The function ValidateIAMPolicy takes an IAMPolicy struct as an argument and returns an error if the policy is not valid.
Each field of the IAMPolicy struct is validated by calling a specific validation function.

Validation is done be checking if required fields are present and if the type and format of each field is correct,
especially for fields defined as interfaces as it is not verified during JSON unmarshalling or Json decoding.
*/

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func ValidateIAMPolicy(policy IAMPolicy) error {
	if _, err := ValidatePolicyName(policy.PolicyName); err != nil {
		return err
	}

	if _, err := ValidatePolicyDocument(policy.PolicyDocument); err != nil {
		return err
	}

	for _, statement := range policy.PolicyDocument.Statements {
		if _, err := ValidateStatement(statement); err != nil {
			return err
		}
	}

	return nil
}

func ValidatePolicyName(name string) (bool, error) {
	if name == "" {
		return false, fmt.Errorf(errorMessages["emptyPolicyName"])
	}

	if reflect.TypeOf(name).Kind() != reflect.String {
		return false, fmt.Errorf(errorMessages["invalidPolicyNameType"])
	}

	matched, err := regexp.MatchString(`^[\w+=,.@-]+$`, name)
	if err != nil {
		return false, fmt.Errorf("error while matching PolicyName: %v", err)
	}

	if !matched {
		return false, fmt.Errorf(errorMessages["invalidPolicyNameFormat"])
	}

	if len(name) > 128 {
		return false, fmt.Errorf(errorMessages["invalidPolicyNameFormat"])
	}
	return true, nil
}

func ValidatePolicyDocument(policyDocument PolicyDocument) (bool, error) {
	if policyDocument.Version == "" {
		return false, fmt.Errorf(errorMessages["invalidPolicyNameFormat"])
	}

	if result, err := ValidateVersion(policyDocument.Version); err != nil {
		return result, err
	}
	return true, nil
}

func ValidateVersion(version interface{}) (bool, error) {
	if reflect.TypeOf(version).Kind() != reflect.String {
		return false, fmt.Errorf(errorMessages["invalidVersionType"])
	}

	versionStr := version.(string)

	if versionStr != "2012-10-17" && versionStr != "2008-10-17" {
		return false, fmt.Errorf(errorMessages["invalidVersionType"])
	}
	return true, nil
}

func ValidateStatement(statement Statement) (bool, error) {
	if result, err := ValidateEffect(statement.Effect); err != nil {
		return result, err
	}
	if len(statement.NotAction) > 0 && len(statement.Action) > 0 {
		return false, fmt.Errorf(errorMessages["emptyActions"])
	}
	if len(statement.NotResource) > 0 && len(statement.Resource) > 0 {
		return false, fmt.Errorf(errorMessages["emptyResource"])
	}
	if len(statement.NotAction) > 0 {
		if result, err := ValidateActions(statement.NotAction); err != nil {
			return result, err
		}
	} else {
		if result, err := ValidateActions(statement.Action); err != nil {
			return result, err
		}
	}
	if len(statement.NotResource) > 0 {
		if result, err := ValidateResources(statement.NotResource); err != nil {
			return result, err
		}
	} else {
		if result, err := ValidateResources(statement.Resource); err != nil {
			return result, err
		}
	}
	return true, nil
}

func ValidateEffect(effect string) (bool, error) {
	if effect == "" {
		return false, fmt.Errorf(errorMessages["emptyEffect"])
	}
	if effect != "Allow" && effect != "Deny" {
		return false, fmt.Errorf(errorMessages["invalidEffect"], effect)
	}
	return true, nil
}

func ValidateActions(actions interface{}) (bool, error) {
	v := reflect.ValueOf(actions)

	if v.Kind() != reflect.Slice {
		return false, fmt.Errorf(errorMessages["invalidActionType"])
	}
	if v.Len() == 0 {
		return false, fmt.Errorf(errorMessages["emptyActions"])
	}
	for i := 0; i < v.Len(); i++ {
		action, ok := v.Index(i).Interface().(string)
		if !ok || !strings.Contains(action, ":") {
			return false, fmt.Errorf(errorMessages["invalidActionFormat"])
		}
	}
	return true, nil
}

func ValidateResources(resources interface{}) (bool, error) {
	v := reflect.ValueOf(resources)

	if v.Kind() != reflect.Slice {
		return false, fmt.Errorf(errorMessages["invalidResourceType"])
	}
	if v.Len() == 0 {
		return false, fmt.Errorf(errorMessages["emptyResources"])
	}

	return true, nil
}
