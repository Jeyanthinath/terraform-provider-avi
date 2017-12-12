/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceClusterNodeRebootSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}