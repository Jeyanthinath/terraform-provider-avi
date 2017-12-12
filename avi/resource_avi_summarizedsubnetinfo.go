/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSummarizedSubnetInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cidr_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
