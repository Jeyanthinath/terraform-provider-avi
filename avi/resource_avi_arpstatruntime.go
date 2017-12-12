/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceArpStatRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dropped": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"dupips": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"proc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"received": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"rxreplies": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"rxrequests": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeouts": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"txreplies": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"txrequests": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
