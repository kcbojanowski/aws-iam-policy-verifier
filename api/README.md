# AWS IAM Policy Verifier API

Simple HTTP service that validates AWS IAM policy documents against the AWS IAM policy JSON structure.

## Getting Started

### Prerequisites
Before you can run the API, make sure you have Go (Golang) installed on your machine. 
You can download and install Go from the [official website](https://golang.org/dl/).

### Running the API
To run the API, follow these steps:
1. Clone the repository:\
`git clone https://github.com/kcbojanowski/remitly-json-verifier`
2. Navigate to the cloned directory:\
`cd aws-iam-policy-verifier/api`
3. Build the project:\
`go build github.com/kcbojanowski/aws-iam-policy-verifier/api`
4. Run the project:\
`./api`
This will start the API server on http://localhost:8080

### Using the API
To validate an IAM policy, make a POST request to `/validate` with a JSON body containing the IAM policy:
```json
{
  "policyName": "ExamplePolicy",
  "policyDocument": {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Action": ["s3:GetObject"],
        "Resource": ["arn:aws:s3:::examplebucket/*"]
      }
    ]
  }
}
```

**Example curl command:**
```
bash curl -X POST -H "Content-Type: application/json" -d @example)policy.json http://localhost:8080/validate
```
Replace your_policy.json with the path to the JSON file containing your IAM policy.

### Response
The API will respond with a JSON object containing the validation result:
* If the policy is valid, the response will be:
```json
{
  "is_valid": true
}
```
* If the policy is invalid, the response will be:
```json
{
  "is_valid": false,
  "error": "Description of the error"
}
```
