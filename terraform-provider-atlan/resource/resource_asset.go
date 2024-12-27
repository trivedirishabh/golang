package atlan

import (
	"github.com/atlanhq/atlan-go/atlan"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAsset() *schema.Resource {
	return &schema.Resource{
		Create: resourceAssetCreate,
		Read:   resourceAssetRead,
		Update: resourceAssetUpdate,
		Delete: resourceAssetDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceAssetCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*atlan.Client)

	asset := models.Asset{
		Name:     d.Get("name").(string),
		Type:     d.Get("type").(string),
		Metadata: d.Get("metadata").(map[string]interface{}),
	}

	createdAsset, err := client.CreateAsset(asset)
	if err != nil {
		return err
	}

	d.SetId(createdAsset.ID)
	return resourceAssetRead(d, m)
}

func resourceAssetRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*atlan.Client)
	assetID := d.Id()

	asset, err := client.GetAsset(assetID)
	if err != nil {
		return err
	}

	d.Set("name", asset.Name)
	d.Set("type", asset.Type)
	d.Set("metadata", asset.Metadata)
	return nil
}

func resourceAssetUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*atlan.Client)
	assetID := d.Id()

	asset := models.Asset{
		Name:     d.Get("name").(string),
		Type:     d.Get("type").(string),
		Metadata: d.Get("metadata").(map[string]interface{}),
	}

	_, err := client.UpdateAsset(assetID, asset)
	if err != nil {
		return err
	}

	return resourceAssetRead(d, m)
}

func resourceAssetDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*atlan.Client)
	assetID := d.Id()

	err := client.DeleteAsset(assetID)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
