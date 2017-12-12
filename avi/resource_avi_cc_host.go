/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCC_HostSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"host_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_last_ts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssh_state": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
