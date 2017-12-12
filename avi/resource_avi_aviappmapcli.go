/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAviAppMapCliSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"nsx_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"proto": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
