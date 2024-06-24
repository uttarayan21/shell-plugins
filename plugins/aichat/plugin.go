package aichat

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "aichat",
		Platform: schema.PlatformInfo{
			Name:     "aichat",
			Homepage: sdk.URL("https://aichat.com"), // TODO: Check if this is correct
		},
		Credentials: []schema.CredentialType{
			APIKey(),
		},
		Executables: []schema.Executable{
			aichatCLI(),
		},
	}
}
