/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceClientLogStreamingConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"external_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  514,
			},
			"log_types_to_send": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LOGS_ALL"},
			"max_logs_per_second": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
		},
	}
}
