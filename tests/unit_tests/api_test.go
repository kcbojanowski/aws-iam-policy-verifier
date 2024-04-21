package unit_tests

import (
	"bytes"
	"encoding/json"
	"github.com/kcbojanowski/aws-iam-policy-verifier/api"
	"github.com/kcbojanowski/aws-iam-policy-verifier/pkg/validator"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

// helper function to generate response from the server
func generateResponse(t *testing.T, server *httptest.Server, filePath, expectedErr string) (api.PolicyResponse, int) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read test file %s: %v", filePath, err)
	}

	resp, err := http.Post(server.URL+"/validate", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatalf("Failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	var gotResponse api.PolicyResponse
	if err := json.NewDecoder(resp.Body).Decode(&gotResponse); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if gotResponse.Error != expectedErr {
		t.Errorf("Expected error message %q; got %q for %s", expectedErr, gotResponse.Error, filePath)
	}

	return gotResponse, resp.StatusCode
}

func TestEmptyFieldsSuite(t *testing.T) {
	var server = httptest.NewServer(http.HandlerFunc(api.ValidateIAMPolicyHandler))
	defer server.Close()

	testCases := []struct {
		filename    string
		expectedErr string
	}{
		{"empty_action.json", validator.GetErrorMessage("emptyAction")},
		{"empty_effect.json", validator.GetErrorMessage("emptyEffect")},
		{"empty_policyname.json", validator.GetErrorMessage("emptyName")},
	}

	for _, tc := range testCases {
		t.Run(tc.filename, func(t *testing.T) {
			filePath := filepath.Join("../test_data/empty_fields", tc.filename)
			gotResponse, statusCode := generateResponse(t, server, filePath, tc.expectedErr)

			if statusCode != http.StatusBadRequest {
				t.Errorf("Expected status code %d; got %d for %s", http.StatusBadRequest, statusCode, tc.filename)
			}

			if gotResponse.IsValid {
				t.Errorf("Expected is_valid to be false; got true for %s", tc.filename)
			}
		})
	}
}

func TestResourceContentSuite(t *testing.T) {
	var server = httptest.NewServer(http.HandlerFunc(api.ValidateIAMPolicyHandler))
	defer server.Close()

	testCases := []struct {
		filename    string
		expectedErr string
	}{
		{"asterisk_resource.json", validator.GetErrorMessage("wildcardResource")},
	}

	for _, tc := range testCases {
		t.Run(tc.filename, func(t *testing.T) {
			filePath := filepath.Join("../test_data/resource_content", tc.filename)
			gotResponse, statusCode := generateResponse(t, server, filePath, tc.expectedErr)

			if statusCode != http.StatusBadRequest {
				t.Errorf("Expected status code %d; got %d for %s", http.StatusBadRequest, statusCode, tc.filename)
			}

			if gotResponse.IsValid {
				t.Errorf("Expected is_valid to be false; got true for %s", tc.filename)
			}
		})
	}
}

func TestValidFormatSuite(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(api.ValidateIAMPolicyHandler))
	defer server.Close()

	files, err := filepath.Glob("tests/test_data/valid_format/*.json")
	if err != nil {
		t.Fatalf("Failed to list test files for valid_format suite: %v", err)
	}

	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			gotResponse, statusCode := generateResponse(t, server, file, "")

			if statusCode != http.StatusOK {
				t.Errorf("Expected status code %d for %s; got %d", http.StatusOK, file, statusCode)
			}

			if !gotResponse.IsValid {
				t.Errorf("Expected is_valid to be true for %s; got false with error: %s", file, gotResponse.Error)
			}

			if gotResponse.Error != "" {
				t.Errorf("Did not expect an error for %s; got error message: %s", file, gotResponse.Error)
			}
		})
	}
}
