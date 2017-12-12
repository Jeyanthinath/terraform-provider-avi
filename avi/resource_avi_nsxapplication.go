/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcensxApplicationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"element": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcensxApplicationElementSchema()},
			"inheritanceallowed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"isuniversal": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nodeid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"objectid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"objecttypename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"revision": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scope": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcensxScopeSchema()},
			"type": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcensxFwObjTypeSchema()},
			"universalrevision": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vsmuuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
