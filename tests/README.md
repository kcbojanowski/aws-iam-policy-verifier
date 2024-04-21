# Tests

# Tests for AWS IAM Role Policy Verifier

This directory contains unit tests for the AWS IAM Role Policy Verifier project. 
The tests are written in Go and use the built-in testing package.

## Structure

`fields_test.go` contains tests for validating individual fields in an IAM policy.
`api_test.go` contains tests for the API endpoint that validates JSON via HTTP POST requests.

## Running the Tests

To run all the tests in this directory, navigate to the `tests` directory in your terminal and run the following command:

```bash
go test ./...