package aichat

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func aichatCLI() schema.Executable {
	return schema.Executable{
		Name:      "aichat CLI", // TODO: Check if this is correct
		Runs:      []string{"aichat"},
		DocsURL:   sdk.URL("https://aichat.com/docs/cli"), // TODO: Replace with actual URL
		NeedsAuth: needsauth.NotForHelpOrVersion(),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.APIKey,
			},
		},
	}
}
