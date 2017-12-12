/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeGeoDbDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"geo_db_profile_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"geo_db_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
