/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceGslbSubDomainPlacementRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"placement_allowed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"sub_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"transition_ops": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_NONE"},
		},
	}
}
