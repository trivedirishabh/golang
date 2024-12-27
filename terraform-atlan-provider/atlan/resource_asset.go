package atlan

import (
	"context"

	"github.com/atlanhq/atlan-go"
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
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the asset",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the asset",
			},
		},
	}
}

func resourceAssetCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*atlan.Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	asset, err := client.Assets.Create(context.Background(), &atlan.Asset{
		Name:        name,
		Description: description,
	})
	if err != nil {
		return err
	}

	d.SetId(asset.GUID)
	return resourceAssetRead(d, m)
}

func resourceAssetRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*atlan.Client)

	asset, err := client.Assets.GetByGUID(context.Background(), d.Id())
	if err != nil {
		return err
	}

	d.Set("name", asset.Name)
	d.Set("description", asset.Description)
	return nil
}

func resourceAssetUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*atlan.Client)

	_, err := client.Assets.Update(context.Background(), &atlan.Asset{
		GUID:        d.Id(),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	})
	if err != nil {
		return err
	}

	return resourceAssetRead(d, m)
}

func resourceAssetDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*atlan.Client)

	err := client.Assets.Delete(context.Background(), d.Id())
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
