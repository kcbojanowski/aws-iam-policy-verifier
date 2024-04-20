# AWS IAM Role Policy JSON Verifier
Project created as a part of the Remitly Poland interview process.

A tool for verifying AWS IAM Role Policy JSON structures. The project provides a simple, interactive command-line interface for validating JSON files against the AWS IAM Role Policy JSON structure.
Additionally, it introduce an endpoint, which can be used to validate JSON via HTTP POST requests.

## All Features
- Validates AWS IAM Role Policy JSON structures
- Provides a CLI for validating JSON files or testing project using internal data
- Provides a web server with an endpoint for validating JSON via HTTP POST requests
- Includes unit tests for all fields in IAM Role Policy JSON structure

## How to Run

To build and run the JSON verifier, follow these steps:

1. Clone the repository:\
`git clone https://github.com/kcbojanowski/remitly-json-verifier`
2. Navigate to cloned directory \
`cd remitly-json-verifier`
3. Build the project\
`go build ./cmd/main.go`
4. Run the project
`./main`

To run web server, use the following command:

1. Go to api directory\
`cd api`
2. Run server.go\
`go run server.go`
3. Server will be available at `http://localhost:8080/validate`

## Resources:
The validation is based on **[Documentation provided by AWS](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html)**

## Testing
To run tests, use the following command:
`go test ./...`
or test in in CLI using test data:\
![img.png](static/internal_data.png)

## Self-Assessment
- [x] Method verifying the input JSON data
- [x] Readme includes "how to run" instructions
- [x] Input data format is defined as AWS::IAM::Role Policy
- [x] Unit tests
- [x] Covering edge cases

