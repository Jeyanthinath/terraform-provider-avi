/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeMgmtIpFailReqSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"reason_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_FAILURE"},
			"subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrPrefixSchema()},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
		},
	}
}
