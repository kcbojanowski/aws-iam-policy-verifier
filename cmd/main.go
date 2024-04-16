package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
    "github.com/manifoldco/promptui"
    "github.com/kcbojanowski/aws-iam-policy-verifier/pkg/verifier"
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
        testdataPath := "./testdata"
        files, err := ioutil.ReadDir(testdataPath)
        if err != nil {
            fmt.Println("Error reading test data directory:", err)
            os.Exit(1)
        }

        for _, f := range files {
            if strings.HasSuffix(f.Name(), ".json") {
                fmt.Printf("\n--- Testing %s:\n", f.Name())
                testJSON, err := ioutil.ReadFile(filepath.Join(testdataPath, f.Name()))
                if err != nil {
                    fmt.Printf("Error reading %s: %s ❌\n", f.Name(), err)
                    continue
                }

                valid, err := verifier.VerifyPolicyJson(testJSON)
                if err != nil {
                    fmt.Printf("Error verifying %s: %s ❌\n", f.Name(), err)
                } else if valid {
                    fmt.Printf("%s is valid. ✅\n", f.Name())
                } else {
                    fmt.Printf("%s is not valid. ❌\n", f.Name())
                }
                fmt.Println()
            }
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

        jsonBytes, err := ioutil.ReadFile(filePath)
        if err != nil {
            fmt.Printf("Error reading JSON file: %s ❌\n", err)
            return
        }

        valid, err := verifier.VerifyPolicyJson(jsonBytes)
        if err != nil {
            fmt.Printf("Error verifying JSON: %s ❌\n", err)
        } else if valid {
            fmt.Println("The JSON policy is valid. ✅")
        } else {
            fmt.Println("The JSON policy is not valid: contains a resource with a single asterisk. ❌")
        }
    }
}
