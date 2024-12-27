package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/trivedirishabh/terraform-provider-atlan/atlan"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: atlan.Provider,
	})
}
