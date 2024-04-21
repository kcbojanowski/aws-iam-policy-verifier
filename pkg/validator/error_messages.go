package validator

var errorMessages = map[string]string{
	"emptyName":         "PolicyName is required and cannot be empty",
	"invalidNameType":   "PolicyName must be String",
	"invalidNameFormat": "PolicyName should match the pattern [\\w+=,.@-]+ and must be <= 128 characters",

	"emptyVersion":       "Version is required and cannot be empty",
	"invalidVersionType": "Version must be a string: '2012-10-17' or '2008-10-17'",

	"emptyEffect":       "Effect is required and cannot be empty",
	"invalidEffect":     "Effect must be 'Allow' or 'Deny'",
	"invalidEffectType": "Effect must be a string",

	"emptyAction":         "At least one Action or NotAction is required",
	"bothActions":         "There can be only one of Action or NotAction",
	"invalidActionFormat": "Invalid action format: each action must include a colon, like 'service:action'",
	"invalidActionType":   "Actions must be a string or a slice of strings",

	"emptyResource":       "At least one Resource or NotResource is required",
	"bothResource":        "There can be only one of Resource or NotResource",
	"invalidResourceType": "Resource must be a string or a slice of strings",
	"wildcardResource":    "Resource is a wildcard",

	"emptyStatement": "At least one Statement is required",
	"unwantedField":  "json: unknown field \"UnwantedField\"",
}

func GetErrorMessage(key string) string {
	return errorMessages[key]
}
