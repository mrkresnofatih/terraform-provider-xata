package resources

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func WorkspaceResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Id of the Xata Workspace",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Name of the Xata Workspace",
			},
			"slug": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "A shorthand for the Xata Workspace name",
			},
		},
	}
}
