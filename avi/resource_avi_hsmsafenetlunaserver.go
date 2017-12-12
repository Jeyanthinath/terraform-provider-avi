/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHSMSafenetLunaServerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"partition_passwd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"partition_serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
