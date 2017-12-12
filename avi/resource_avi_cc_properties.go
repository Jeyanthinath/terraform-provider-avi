/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCC_PropertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rpc_poll_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"rpc_queue_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
		},
	}
}
