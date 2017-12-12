/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSamlServiceProviderSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"org_display_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"org_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"org_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"saml_entity_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AUTH_SAML_CLUSTER_VIP"},
			"sp_nodes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSamlServiceProviderNodeSchema()},
			"tech_contact_email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tech_contact_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
