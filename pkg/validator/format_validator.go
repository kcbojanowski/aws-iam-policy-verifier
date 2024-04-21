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

	if len(policy.PolicyDocument.Statement) == 0 || policy.PolicyDocument.Statement == nil {
		return fmt.Errorf(errorMessages["emptyStatement"])
	}

	for _, statement := range policy.PolicyDocument.Statement {
		if _, err := ValidateStatement(statement); err != nil {
			return err
		}
	}

	return nil
}

func ValidatePolicyDocument(policyDocument PolicyDocument) (bool, error) {
	if policyDocument.Version == "" {
		return false, fmt.Errorf(errorMessages["emptyVersion"])
	}

	if result, err := ValidateVersion(policyDocument.Version); err != nil {
		return result, err
	}
	return true, nil
}

func ValidateStatement(statement Statement) (bool, error) {
	if result, err := ValidateEffect(statement.Effect); err != nil {
		return result, err
	}

	if (statement.Action != nil) && (statement.NotAction != nil) {
		return false, fmt.Errorf(errorMessages["bothActions"])
	}

	if (statement.Resource != nil) && (statement.NotResource != nil) {
		return false, fmt.Errorf(errorMessages["bothResources"])
	}

	if statement.Action != nil {
		if result, err := ValidateActions(statement.Action); !result {
			return result, err
		}
	}

	if statement.NotAction != nil {
		if result, err := ValidateActions(statement.NotAction); !result {
			return result, err
		}
	}

	if statement.Resource != nil {
		if result, err := ValidateResources(statement.Resource); !result {
			return result, err
		}
	}

	if statement.NotResource != nil {
		if result, err := ValidateResources(statement.NotResource); !result {
			return result, err
		}
	}

	return true, nil
}

func ValidatePolicyName(name string) (bool, error) {
	if name == "" {
		return false, fmt.Errorf(errorMessages["emptyName"])
	} else {
		if reflect.TypeOf(name).Kind() != reflect.String {
			return false, fmt.Errorf(errorMessages["invalidNameType"])
		}

		matched, err := regexp.MatchString(`^[\w+=,.@-]+$`, name)
		if err != nil {
			return false, fmt.Errorf("error while matching PolicyName: %v", err)
		}

		if !matched {
			return false, fmt.Errorf(errorMessages["invalidNameFormat"])
		}

		if len(name) > 128 {
			return false, fmt.Errorf(errorMessages["invalidNameFormat"])
		}

		return true, nil
	}
}

func ValidateVersion(version interface{}) (bool, error) {
	versionStr, ok := version.(string)

	if !ok || versionStr == "" {
		return false, fmt.Errorf(errorMessages["emptyVersion"])
	}

	if reflect.TypeOf(version).Kind() != reflect.String {
		return false, fmt.Errorf(errorMessages["invalidVersionType"])
	}

	if versionStr != "2012-10-17" && versionStr != "2008-10-17" {
		return false, fmt.Errorf(errorMessages["invalidVersionType"])
	}
	return true, nil
}

func ValidateActions(actions interface{}) (bool, error) {
	v := reflect.ValueOf(actions)
	if !v.IsValid() {
		return false, fmt.Errorf(errorMessages["emptyAction"])
	}

	var actionsList []string
	switch v.Kind() {
	case reflect.String:
		actionsList = append(actionsList, v.String())
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i)
			if str, ok := item.Interface().(string); ok {
				actionsList = append(actionsList, str)
			} else {
				return false, fmt.Errorf(errorMessages["invalidActionType"])
			}
		}
	default:
		return false, fmt.Errorf(errorMessages["invalidActionType"])
	}

	if len(actionsList) == 0 {
		return false, fmt.Errorf(errorMessages["emptyAction"])
	}

	for _, action := range actionsList {
		if !strings.Contains(action, ":") {
			return false, fmt.Errorf(errorMessages["invalidActionFormat"])
		}
	}
	return true, nil
}

func ValidateResources(resources interface{}) (bool, error) {
	v := reflect.ValueOf(resources)

	if !v.IsValid() {
		return false, fmt.Errorf(errorMessages["invalidResourceType"])
	}

	switch v.Kind() {
	case reflect.String:
		str := v.String()
		if str == "" {
			return false, fmt.Errorf(errorMessages["emptyResource"])
		}
		if str == "*" {
			return false, fmt.Errorf(errorMessages["wildcardResource"])
		}
	case reflect.Slice:
		if v.Len() == 0 {
			return false, fmt.Errorf(errorMessages["emptyResource"])
		}
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i)
			if !item.CanInterface() {
				return false, fmt.Errorf(errorMessages["invalidResourceFormat"])
			}
			str, ok := item.Interface().(string)
			if !ok {
				return false, fmt.Errorf(errorMessages["invalidResourceType"])
			}
			if str == "*" {
				return false, fmt.Errorf(errorMessages["wildcardResource"])
			}
		}
	default:
		return false, fmt.Errorf(errorMessages["invalidResourceType"])
	}

	return true, nil
}

func ValidateEffect(effect string) (bool, error) {
	if effect == "" {
		return false, fmt.Errorf(errorMessages["emptyEffect"])
	}
	if effect != "Allow" && effect != "Deny" {
		return false, fmt.Errorf(errorMessages["invalidEffect"])
	}
	return true, nil
}
