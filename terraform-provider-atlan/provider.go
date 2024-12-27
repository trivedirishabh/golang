package atlan

import (
	"github.com/atlanhq/atlan-go/atlan" // Import the Atlan SDK
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		// Define provider schema for configuration (API Key, etc.)
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ATL_API_KEY", nil),
				Description: "API key for Atlan authentication",
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ATL_BASE_URL", "https://api.atlan.com"),
				Description: "Base URL for the Atlan API",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"atlan_asset": resourceAsset(),
		},
		ConfigureFunc: providerConfigure,
	}
}

// providerConfigure sets up the client for use in resources
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	apiKey := d.Get("api_key").(string)
	baseURL := d.Get("base_url").(string)

	client, err := atlan.NewClient(apiKey, baseURL)
	if err != nil {
		return nil, err
	}
	return client, nil
}
