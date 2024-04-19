package validator

import (
	"bytes"
	"encoding/json"
)

type IAMPolicy struct {
	PolicyId       string         `json:"Id,omitempty"`
	PolicyName     string         `json:"PolicyName"`
	PolicyDocument PolicyDocument `json:"PolicyDocument"`
}

type PolicyDocument struct {
	Version    string      `json:"Version"`
	Statements []Statement `json:"Statement"`
}

type Statement struct {
	Sid         string   `json:"Sid,omitempty"`
	Effect      string   `json:"Effect"`
	Action      []string `json:"Action,omitempty"`
	NotAction   []string `json:"NotAction,omitempty"`
	Resource    []string `json:"Resource"`
	NotResource []string `json:"NotResource,omitempty"`
}

func loadPolicyFromJSON(data []byte) (IAMPolicy, error) {
	var policy IAMPolicy
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&policy)
	if err != nil {
		return IAMPolicy{}, err
	}
	return policy, nil
}
