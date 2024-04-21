package main

import (
	"fmt"
	"github.com/kcbojanowski/aws-iam-policy-verifier/api"
	"github.com/kcbojanowski/aws-iam-policy-verifier/pkg/validator"
	"github.com/manifoldco/promptui"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
)

func main() {
	mode, err := selectMode()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch mode {
	case "Test with internal data":
		validateInternalData("./tests/test_data")
	case "Input your own JSON file":
		validateUserFile()
	case "Run Server":
		runServer()
	}
}

func selectMode() (string, error) {
	prompt := promptui.Select{
		Label: "Select Mode",
		Items: []string{"Test with internal data", "Input your own JSON file", "Run Server"},
	}
	_, result, err := prompt.Run()
	return result, err
}

func validateInternalData(testdataPath string) {
	err := filepath.Walk(testdataPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error accessing path:", path, "error:", err)
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			fmt.Printf("\n--- Testing %s:\n", info.Name())
			validateFile(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking through test data directory:", err)
	}
}

func validateUserFile() {
	prompt := promptui.Prompt{
		Label: "Enter path to the JSON file",
	}
	filePath, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	validateFile(filePath)
}

func validateFile(filePath string) {
	valid, err := validator.ValidatePolicyJson(filePath)
	if err != nil || !valid {
		fmt.Printf("Validation failed for %s: %s\n", filePath, err)
	} else {
		fmt.Printf("Validation successful for %s\n", filePath)
	}
}

func runServer() {
	fmt.Println("Starting server on http://localhost:8080...")
	http.HandleFunc("/validate", api.ValidateIAMPolicyHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
