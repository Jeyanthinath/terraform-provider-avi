/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceIpamDnsCustomProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"custom_ipam_dns_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_params": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCustomParamsSchema()},
			"usable_domains": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"usable_subnets": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema()},
		},
	}
}
