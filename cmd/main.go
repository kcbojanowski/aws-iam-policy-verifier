package main

import (
	"fmt"
	"github.com/kcbojanowski/aws-iam-policy-verifier/pkg/validator"
	"github.com/manifoldco/promptui"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	prompt := promptui.Select{
		Label: "Select Mode",
		Items: []string{"Test with internal data", "Input your own JSON file"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "Test with internal data":
		testdataPath := "./test-data"
		err := filepath.Walk(testdataPath, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				fmt.Println("Error accessing path:", path, "error:", err)
				return err
			}
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
				fmt.Printf("\n--- Testing %s:\n", info.Name())
				return validator.ValidatePolicyJson(path)
			}
			return nil
		})
		if err != nil {
			fmt.Println("Error walking through test data directory:", err)
			os.Exit(1)
		}

	case "Input your own JSON file":
		prompt := promptui.Prompt{
			Label: "Enter path to the JSON file",
		}
		filePath, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		if err := validateJSON(filePath); err != nil {
			fmt.Printf("Validation failed for %s: %s\n", filePath, err)
		}
	}
}
