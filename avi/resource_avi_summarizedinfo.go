/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSummarizedInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subnet_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSummarizedSubnetInfoSchema()},
		},
	}
}
