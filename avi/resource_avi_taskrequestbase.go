/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceTaskRequestBaseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"callback_queue": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"method_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"request_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"transaction_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
