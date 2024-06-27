package aichat

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/importer"
	"github.com/1Password/shell-plugins/sdk/provision"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func APIKey() schema.CredentialType {
	return schema.CredentialType{
		Name:          credname.APIKey,
		DocsURL:       sdk.URL("https://github.com/sigoden/aichat"),
		ManagementURL: sdk.URL("https://github.com/sigoden/aichat"),
		Fields: []schema.CredentialField{
			{
				Name:                fieldname.APIKey,
				MarkdownDescription: "API Key used to authenticate to openai.",
				Secret:              true,
			},
			{
				Name:                fieldname.AppSecret,
				MarkdownDescription: "Platform used to authenticate to aichat.",
				Secret:              false,
			},
		},
		DefaultProvisioner: provision.EnvVars(defaultEnvVarMapping),
		Importer:           importer.TryEnvVarPair(defaultEnvVarMapping),
	}
}

var defaultEnvVarMapping = map[string]sdk.FieldName{
	"OPENAI_API_KEY":  fieldname.APIKey,
	"AICHAT_PLATFORM": fieldname.AppSecret,
}
