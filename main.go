package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/atlassian/terraform-provider-artifactory/pkg/artifactory"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: artifactory.Provider,
	})
}
