/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceFiuConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"fiu_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "FIU_ACTION_NONE"},
			"fiu_filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceFiuFilterSchema()},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"probability": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
		},
	}
}
