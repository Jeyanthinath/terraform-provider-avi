/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCloudTenantsDeletedSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenants": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCloudTenantCleanupSchema()},
			"vtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
