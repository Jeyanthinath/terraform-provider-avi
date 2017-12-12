/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHostInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu_ok": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"disk_ok": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"host_resource_scale": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mem_ok": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"mgmt_nw_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"num_se_possible": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reason_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_ok": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"se_mgmt_ok": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
		},
	}
}
