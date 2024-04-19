package validator

var errorMessages = map[string]string{
	"emptyPolicyName":         "PolicyName is required and cannot be empty",
	"invalidPolicyNameType":   "PolicyName must be String",
	"invalidPolicyNameFormat": "PolicyName should match the pattern [\\w+=,.@-]+ and must be <= 128 characters",

	"emptyVersion":       "Version is required and cannot be empty",
	"invalidVersionType": "Version must be a string: '2012-10-17' or '2008-10-17'",

	"emptyEffect":       "Effect is required and cannot be empty",
	"invalidEffect":     "Effect must be 'Allow' or 'Deny'",
	"invalidEffectType": "Effect must be a string",

	"emptyAction":         "At least one Action or NotAction is required",
	"invalidActionFormat": "Invalid action format: each action must include a colon, like 'service:action'",
	"invalidActionType":   "Actions must be a slice of strings",

	"emptyResources":      "At least one Resource or NotResource is required",
	"invalidResourceType": "Resources must be a slice of strings",

	"emptyStatements":      "At least one statement is required in PolicyDocument",
	"invalidStatementType": "Statements must be a slice of Statement structs",
}
