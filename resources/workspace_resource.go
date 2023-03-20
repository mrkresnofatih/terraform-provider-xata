package resources

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mrkresnofatih/terraform-provider-xata/services/base"
	"github.com/mrkresnofatih/terraform-provider-xata/services/workspaceservice"
)

func WorkspaceResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Id of the Xata Workspace",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Name of the Xata Workspace",
			},
			"slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A shorthand for the Xata Workspace name",
			},
		},
		CreateContext: func(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
			var diags diag.Diagnostics
			client := i.(base.XataApi)

			name := data.Get("name").(string)
			slug := data.Get("slug").(string)

			svc := workspaceservice.WorkspaceService{
				XataApi: client,
			}

			createdWorkspace, err := svc.Create(workspaceservice.WorkspaceCreateRequest{
				Name: name,
				Slug: slug,
			})
			if err != nil {
				return diag.FromErr(err)
			}
			if err = data.Set("name", createdWorkspace.Name); err != nil {
				return diag.FromErr(err)
			}
			if err = data.Set("slug", createdWorkspace.Slug); err != nil {
				return diag.FromErr(err)
			}
			data.SetId(createdWorkspace.Id)
			return diags
		},
		ReadContext: func(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
			var diags diag.Diagnostics
			client := i.(base.XataApi)
			id := data.Get("id").(string)
			svc := workspaceservice.WorkspaceService{
				XataApi: client,
			}
			foundWorkspace, err := svc.Get(workspaceservice.WorkspaceGetRequest{
				Id: id,
			})
			if err != nil {
				return diag.FromErr(err)
			}
			if err = data.Set("name", foundWorkspace.Name); err != nil {
				return diag.FromErr(err)
			}
			if err = data.Set("slug", foundWorkspace.Slug); err != nil {
				return diag.FromErr(err)
			}
			data.SetId(foundWorkspace.Id)
			return diags
		},
	}
}
