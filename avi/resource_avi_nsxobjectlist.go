/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcensxObjectListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcensxApplicationSchema()},
			"ipset": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceipsetSchema()},
			"securitygroup": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcesecuritygroupSchema()},
		},
	}
}
