package unit_tests

import (
	"github.com/kcbojanowski/aws-iam-policy-verifier/pkg/validator"
	"testing"
)

func TestValidateResources(t *testing.T) {
	tests := []struct {
		name     string
		input    []interface{}
		expected bool
		errMsg   string
	}{
		{
			name:     "Empty Resource",
			input:    []interface{}{},
			expected: false,
			errMsg:   validator.GetErrorMessage("emptyResource"),
		},
		{
			name:     "Valid Resource",
			input:    []interface{}{"arn:aws:s3:::my_corporate_bucket/*"},
			expected: true,
			errMsg:   "",
		},
		{
			name:     "Invalid Resource Type",
			input:    []interface{}{123},
			expected: false,
			errMsg:   validator.GetErrorMessage("invalidResourceType"),
		},
		{
			name:     "Wildcard Resource",
			input:    []interface{}{"*"},
			expected: false,
			errMsg:   validator.GetErrorMessage("wildcardResource"),
		},
		{
			name:     "Multiple Asterisks",
			input:    []interface{}{"**"},
			expected: true,
			errMsg:   "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := validator.ValidateResources(test.input)

			checkTestResult(t, test.name, test.expected, test.errMsg, result, err)
		})
	}
}

func TestValidatePolicyName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
		errMsg   string
	}{
		{
			name:     "Empty Name",
			input:    "",
			expected: false,
			errMsg:   validator.GetErrorMessage("emptyName"),
		},
		{
			name:     "Valid Name",
			input:    "valid-policy-1",
			expected: true,
			errMsg:   "",
		},
		{
			name:     "Invalid Characters",
			input:    "invalid!name",
			expected: false,
			errMsg:   validator.GetErrorMessage("invalidNameFormat"),
		},
		{
			name:     "Too Long Name",
			input:    repeatString("a", 130),
			expected: false,
			errMsg:   validator.GetErrorMessage("invalidNameFormat"),
		},
		{
			name:     "Valid Complex Name",
			input:    "valid_policy.name=1,admin@org-company",
			expected: true,
			errMsg:   "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := validator.ValidatePolicyName(test.input)

			checkTestResult(t, test.name, test.expected, test.errMsg, result, err)
		})
	}
}

func TestValidateEffect(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
		errMsg   string
	}{
		{
			name:     "Empty Effect",
			input:    "",
			expected: false,
			errMsg:   validator.GetErrorMessage("emptyEffect"),
		},
		{
			name:     "Valid Effect - Allow",
			input:    "Allow",
			expected: true,
			errMsg:   "",
		},
		{
			name:     "Valid Effect - Deny",
			input:    "Deny",
			expected: true,
			errMsg:   "",
		},
		{
			name:     "Invalid Effect",
			input:    "Invalid",
			expected: false,
			errMsg:   validator.GetErrorMessage("invalidEffect"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := validator.ValidateEffect(test.input)

			checkTestResult(t, test.name, test.expected, test.errMsg, result, err)
		})
	}
}

func TestValidatePolicyDocument(t *testing.T) {
	tests := []struct {
		name     string
		input    validator.PolicyDocument
		expected bool
		errMsg   string
	}{
		{
			name: "Valid Policy Document",
			input: validator.PolicyDocument{
				Version: "2012-10-17",
				Statement: []validator.Statement{
					{
						Effect:   "Allow",
						Action:   []string{"s3:GetObject"},
						Resource: []string{"arn:aws:s3:::my_corporate_bucket/*"},
					},
				},
			},
			expected: true,
			errMsg:   "",
		},
		{
			name:     "Empty Policy Document",
			input:    validator.PolicyDocument{},
			expected: false,
			errMsg:   validator.GetErrorMessage("emptyVersion"),
		},
		{
			name: "Invalid Version",
			input: validator.PolicyDocument{
				Version: "Invalid",
				Statement: []validator.Statement{
					{
						Effect:   "Allow",
						Action:   []string{"s3:GetObject"},
						Resource: []string{"arn:aws:s3:::my_corporate_bucket/*"},
					},
				},
			},
			expected: false,
			errMsg:   validator.GetErrorMessage("invalidVersionType"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := validator.ValidatePolicyDocument(test.input)
			checkTestResult(t, test.name, test.expected, test.errMsg, result, err)
		})
	}
}

func TestValidateActions(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected bool
		errMsg   string
	}{
		{
			name:     "Empty Actions",
			input:    []interface{}{},
			expected: false,
			errMsg:   validator.GetErrorMessage("emptyAction"),
		},
		{
			name:     "Valid Actions",
			input:    []interface{}{"s3:PutObject"},
			expected: true,
			errMsg:   "",
		},
		{
			name:     "Invalid Actions Format",
			input:    "no-colon-included",
			expected: false,
			errMsg:   validator.GetErrorMessage("invalidActionFormat"),
		},
		{
			name:     "Valid NotActions",
			input:    []interface{}{"s3:PutObject"},
			expected: true,
			errMsg:   "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := validator.ValidateActions(tc.input)
			checkTestResult(t, tc.name, tc.expected, tc.errMsg, result, err)
		})
	}
}

func TestValidateStatement(t *testing.T) {
	tests := []struct {
		name      string
		statement validator.Statement
		expected  bool
		errMsg    string
	}{
		{
			name:      "Empty Statement",
			statement: validator.Statement{},
			expected:  false,
			errMsg:    validator.GetErrorMessage("emptyEffect"),
		},
		{
			name: "Valid Statement",
			statement: validator.Statement{
				Effect:   "Allow",
				Action:   []interface{}{"s3:PutObject"},
				Resource: []interface{}{"arn:aws:s3:::my_corporate_bucket/*"},
			},
			expected: true,
			errMsg:   "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := validator.ValidateStatement(tc.statement)
			checkTestResult(t, tc.name, tc.expected, tc.errMsg, result, err)
		})
	}
}

/*
Helper functions
*/
func checkTestResult(t *testing.T, name string, expected bool, errMsg string, result bool, err error) {
	if result != expected {
		t.Errorf("%s: expected result %v, got %v", name, expected, result)
	}
	if err != nil && errMsg == "" {
		t.Errorf("%s: expected no error, got %v", name, err)
	} else if err == nil && errMsg != "" {
		t.Errorf("%s: expected error msg '%s', got no error", name, errMsg)
	} else if err != nil && err.Error() != errMsg {
		t.Errorf("%s: expected error msg '%s', got '%v'", name, errMsg, err)
	}
}

func repeatString(s string, n int) string {
	var repeatedString string
	for i := 0; i < n; i++ {
		repeatedString += s
	}
	return repeatedString
}
