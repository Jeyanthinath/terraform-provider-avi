/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVITenantNetworkMapSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"tenant_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf_context_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
