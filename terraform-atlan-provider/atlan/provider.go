package atlan

import (
	"github.com/atlanhq/atlan-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider defines the Terraform provider for Atlan.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "API key for Atlan",
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "https://api.atlan.com",
				Description: "Base URL for the Atlan API",
			},
			"atlan_asset": resourceAsset(),
		},
		ConfigureFunc: providerConfigure,
	}
}

// providerConfigure sets up the client for Atlan API communication.
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	apiKey := d.Get("api_key").(string)
	baseURL := d.Get("base_url").(string)

	client, err := atlan.NewClient(atlan.ClientConfig{
		APIKey:  apiKey,
		BaseURL: baseURL,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
