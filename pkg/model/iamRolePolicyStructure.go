package model

type IAMPolicy struct {
    PolicyName     string         `json:"PolicyName"`
    PolicyDocument PolicyDocument `json:"PolicyDocument"`
}

type PolicyDocument struct {
    Version   string      `json:"Version"`
    Statement []Statement `json:"Statement"`
}

type Statement struct {
    Sid      string   `json:"Sid"`
    Effect   string   `json:"Effect"`
    Action   []string `json:"Action"`
    Resource string   `json:"Resource"`
}
