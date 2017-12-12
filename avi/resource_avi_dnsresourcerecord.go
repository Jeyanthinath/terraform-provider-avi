/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDnsResourceRecordSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"addr_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dclass": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceGeoLocationSchema()},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"site_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"vs_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
