/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCC_DiscoverySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"discovery_tenants": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCC_TenantSchema()},
			"failed_tenants": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCC_TenantSchema()},
			"poll": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCC_CronSchema()},
			"ran_discovery": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"tracked_tenants": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCC_TenantSchema()},
		},
	}
}
