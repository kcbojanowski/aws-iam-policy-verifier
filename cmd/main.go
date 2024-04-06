package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "github.com/kcbojanowski/remitly-json-verifier/pkg/verifier"
    "github.com/manifoldco/promptui"
)

func main() {
    prompt := promptui.Prompt{
        Label: "Enter path to the JSON file",
        Validate: func(input string) error {
            if input == "" {
                return fmt.Errorf("file path cannot be empty")
            }
            if _, err := os.Stat(input); os.IsNotExist(err) {
                return fmt.Errorf("file does not exist")
            }
            return nil
        },
    }

    filePath, err := prompt.Run()
    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        return
    }

    jsonBytes, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Printf("Error reading JSON file: %s\n", err)
        return
    }

    valid, err := verifier.VerifyPolicyJSON(jsonBytes)
    if err != nil {
        fmt.Printf("Error verifying JSON: %s\n", err)
        return
    }

    if valid {
        fmt.Println("The JSON policy is valid.")
    } else {
        fmt.Println("The JSON policy is not valid: contains a resource with a single asterisk.")
    }
}
