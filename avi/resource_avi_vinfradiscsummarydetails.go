/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVinfraDiscSummaryDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"num_clusters": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_dcs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_hosts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_nws": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_vms": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vcenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
