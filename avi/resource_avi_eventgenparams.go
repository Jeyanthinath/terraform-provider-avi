/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceEventGenParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"events": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"pool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sslkeyandcertificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtualservice": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
