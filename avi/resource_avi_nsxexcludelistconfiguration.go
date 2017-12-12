/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcensxExcludeListConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"clienthandle": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"excludemember": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcensxExcludeMemberSchema()},
			"isuniversal": &schema.Schema{
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
			"type": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcensxFwObjTypeSchema()},
			"universalrevision": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vmsuuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
