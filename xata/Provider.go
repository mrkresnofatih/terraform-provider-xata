package xata

import (
	"context"
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mrkresnofatih/terraform-provider-xata/resources"
	"github.com/mrkresnofatih/terraform-provider-xata/services/base"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"xata_workspace": resources.WorkspaceResource(),
		},
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Token for the Authorization Method",
			},
		},
		ConfigureContextFunc: func(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
			var diags diag.Diagnostics
			apiKey := data.Get("api_key").(string)
			if len(apiKey) == 0 {
				return nil, diag.FromErr(errors.New("api key is empty"))
			}
			return base.XataApi{
				ApiKey: apiKey,
			}, diags
		},
	}
}
