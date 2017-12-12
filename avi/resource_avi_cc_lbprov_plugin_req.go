/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Resourcecc_lbprov_plugin_reqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "cloud-0"},
			"form": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceOsLbProvPluginFormSchema()},
		},
	}
}
