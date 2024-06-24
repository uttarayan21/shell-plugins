package aichat

import (
	"context"

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
		},
		DefaultProvisioner: provision.EnvVars(defaultEnvVarMapping),
		Importer: importer.TryAll(
			importer.TryEnvVarPair(defaultEnvVarMapping),
			TryaichatConfigFile(),
		)}
}

var defaultEnvVarMapping = map[string]sdk.FieldName{
	"OPENAI_API_KEY":  fieldname.APIKey,
	"AICHAT_PLATFORM": "openai",
}

// TODO: Check if the platform stores the API Key in a local config file, and if so,
// implement the function below to add support for importing it.
func TryaichatConfigFile() sdk.Importer {
	return importer.TryFile("~/.config/aichat/config.yaml", func(ctx context.Context, contents importer.FileContents, in sdk.ImportInput, out *sdk.ImportAttempt) {
		var config Config
		if err := contents.ToYAML(&config); err != nil {
			out.AddError(err)
			return
		}

		if config.APIKey == "" {
			return
		}

		// if config.Platform == "" {
		// 	return
		// }

		out.AddCandidate(sdk.ImportCandidate{
			Fields: map[sdk.FieldName]string{
				fieldname.APIKey: config.APIKey,
			},
		})
	})
}

// TODO: Implement the config file schema
type Config struct {
	APIKey string
	// Platform string
}
