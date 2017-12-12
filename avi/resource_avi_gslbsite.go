/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceGslbSiteSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_vses": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbSiteDnsVsSchema()},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"hm_proxies": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbHealthMonitorProxySchema()},
			"ip_addresses": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema()},
			"location": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceGslbGeoLocationSchema()},
			"member_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_PASSIVE_MEMBER"},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  443,
			},
			"ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
