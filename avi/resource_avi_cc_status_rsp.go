/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Resourcecc_status_rspSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"agents": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCloudStatusSchema()},
			"ret_status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"ret_string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
