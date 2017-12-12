/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceIpamDnsInfobloxProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_view": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default"},
			"ip_address": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"network_view": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default"},
			"usable_domains": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"usable_subnets": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema()},
			"wapi_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "2.0"},
		},
	}
}
