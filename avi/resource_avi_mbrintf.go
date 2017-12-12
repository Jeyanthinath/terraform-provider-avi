/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMbrIntfSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"if_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"linux_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
