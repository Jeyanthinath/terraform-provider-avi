/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceEtherHeaderSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ether_dhost": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ether_shost": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ether_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
