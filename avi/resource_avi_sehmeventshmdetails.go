/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeHmEventShmDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"average_response_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"health_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"resp_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"response_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
