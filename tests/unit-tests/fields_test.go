package unit_tests

import (
	"github.com/kcbojanowski/aws-iam-policy-verifier/pkg/validator"
	"testing"
)

type testCase struct {
	name     string
	input    string
	expected bool
	errMsg   string
}

func TestValidateResources(t *testing.T) {
}

func TestValidatePolicyName(t *testing.T) {
	tests := []testCase{
		{"Empty Name", "", false, "Policy name cannot be empty"},
		{"Valid Name", "valid-policy-1", true, ""},
		{"Invalid Characters", "invalid!name", false, "Policy name format is invalid"},
		{"Too Long Name", "a" + repeat("a", 128), false, "Policy name format is invalid"},
		{"Valid Complex Name", "valid_policy.name=1,admin@org-company", true, ""},
	}

	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			result, err := validator.ValidatePolicyName(scenario.input)

			if (err != nil) != (scenario.errMsg != "") || (err != nil && err.Error() != scenario.errMsg) {
				t.Errorf("For %s, expected error msg: '%s', actual: '%v'", scenario.name, scenario.errMsg, err)
			}

			if result != scenario.expected {
				t.Errorf("For %s, expected result %v, actual: %v", scenario.name, scenario.expected, result)
			}
		})
	}
}

func TestValidateEffect(t *testing.T) {
	tests := []testCase{
		{"Empty Effect", "", false, "Effect is required and cannot be empty"},
		{"Valid Effect - Allow", "Allow", true, ""},
		{"Valid Effect - Deny", "Deny", true, ""},
		{"Invalid Effect", "invalid-effect", false, "Effect must be 'Allow' or 'Deny'"},
	}

	for _, scenario := range tests {
		t.Run(scenario.name, func(t *testing.T) {
			result, err := validator.ValidateEffect(scenario.input)

			if (err != nil) != (scenario.errMsg != "") || (err != nil && err.Error() != scenario.errMsg) {
				t.Errorf("For %s, expected error msg: '%s', actual: '%v'", scenario.name, scenario.errMsg, err)
			}

			if result != scenario.expected {
				t.Errorf("For %s, expected result %v, actual: %v", scenario.name, scenario.expected, result)
			}
		})
	}
}

func TestValidatePolicyDocument(t *testing.T) {

}

func TestValidateActions(t *testing.T) {
}

func TestValidateStatement(t *testing.T) {
}

func TestValidateIAMPolicy(t *testing.T) {
}

// test utils
func repeat(s string, count int) string {
	str := ""
	for i := 0; i < count; i++ {
		str += s
	}
	return str
}
