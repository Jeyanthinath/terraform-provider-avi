/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceInpcbProxyInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"inp_proxy_ip_minttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inp_proxy_ip_p": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inp_proxy_ip_tos": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inp_proxy_ip_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
