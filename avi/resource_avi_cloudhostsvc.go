/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCloudHostSvcSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"available": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
