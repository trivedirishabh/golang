package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/trivedirishabh/terraform-provider-atlan/provider" // Import the provider package
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider, // Call the provider from the provider package
	})
}
