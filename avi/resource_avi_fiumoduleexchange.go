/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceFiuModuleExchangeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"controller": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "fiu_mod_controller_exchange"},
			"dp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "fiu_mod_dataplane_exchange"},
			"seagent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "fiu_mod_seagent_exchange"},
		},
	}
}
