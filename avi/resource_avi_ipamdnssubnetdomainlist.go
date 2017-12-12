/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceIpamDnsSubnetDomainListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_records": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpamDnsRecordInfoSchema()},
			"domains": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"error": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"subnets": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrPrefixSchema()},
			"zdomains": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsDomainSchema()},
		},
	}
}
