package main

import (
	atlan "terraform-provider-atlan"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: atlan.Provider,
	})
}
