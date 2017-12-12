/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceClusterNodeStateSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"up_since": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
