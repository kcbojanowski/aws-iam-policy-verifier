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
	Version   string      `json:"Version"`
	Statement []Statement `json:"Statement"`
}

type Statement struct {
	Sid          string                  `json:"Sid,omitempty"`
	Principal    *PrincipalBlock         `json:"Principal,omitempty"`
	NotPrincipal *PrincipalBlock         `json:"NotPrincipal,omitempty"`
	Effect       string                  `json:"Effect"`
	Action       interface{}             `json:"Action,omitempty"`
	NotAction    interface{}             `json:"NotAction,omitempty"`
	Resource     interface{}             `json:"Resource"`
	NotResource  interface{}             `json:"NotResource,omitempty"`
	Conditions   map[string]ConditionMap `json:"Conditions,omitempty"`
}

type PrincipalBlock struct {
	AWS           interface{} `json:"AWS,omitempty" validate:"optional"`
	Federated     interface{} `json:"Federated,omitempty" validate:"optional"`
	Service       interface{} `json:"Service,omitempty" validate:"optional"`
	CanonicalUser interface{} `json:"CanonicalUser,omitempty" validate:"optional"`
}

type ConditionMap map[string][]string

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
