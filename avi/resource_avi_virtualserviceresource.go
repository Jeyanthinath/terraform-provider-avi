/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVirtualServiceResourceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_exclusive": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"memory": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_standby_se": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vcpus": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scalein_primary": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"scalein_se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
