package model

type IAMPolicy struct {
    PolicyId       string         `json:"Id,omitempty" validate:"optional,uuid"`
    PolicyName     string         `json:"PolicyName" validate:"required"`
    PolicyDocument PolicyDocument `json:"PolicyDocument" validate:"required"`
}

type PolicyDocument struct {
    Version   string      `json:"Version" validate:"required,version"`
    Statements []Statement `json:"Statement" validate:"required,min=1"`
}

type Statement struct {
    Sid         string         `json:"Sid,omitempty" validate:"optional"`
    Principal   *PrincipalBlock `json:"Principal,omitempty" validate:"optional"`
    NotPrincipal *PrincipalBlock `json:"NotPrincipal,omitempty" validate:"optional"`
    Effect      string         `json:"Effect" validate:"required,effect"`
    Action      []string       `json:"Action,omitempty" validate:"required,min=1"`
    NotAction   []string       `json:"NotAction,omitempty" validate:"optional"`
    Resource    []string       `json:"Resource,omitempty" validate:"required,min=1"`
    NotResource []string       `json:"NotResource,omitempty" validate:"optional"`
    Conditions  map[string]ConditionMap `json:"Condition,omitempty" validate:"optional"`
}

type PrincipalBlock struct {
    AWS          []string `json:"AWS,omitempty" validate:"optional"`
    Federated    []string `json:"Federated,omitempty" validate:"optional"`
    Service      []string `json:"Service,omitempty" validate:"optional"`
    CanonicalUser []string `json:"CanonicalUser,omitempty" validate:"optional"`
}

type ConditionMap map[string][]string

