/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVIDeleteSEReqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"segroup_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sevm_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"vcenter_admin": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceVIAdminCredentialsSchema()},
		},
	}
}
