/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceGslbHealthMonitorProxySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"proxy_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_HEALTH_MONITOR_PROXY_PRIVATE_MEMBERS"},
			"site_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
