/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Resourcecc_webhook_mesos_reqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "cloud-0"},
			"response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"response_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"status_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
