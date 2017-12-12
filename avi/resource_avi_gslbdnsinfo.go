/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceGslbDnsInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_active": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dns_vs_states": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPerDnsStateSchema()},
			"gs_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceGslbDnsGsStatusSchema()},
			"retry_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_table": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbDnsSeInfoSchema()},
		},
	}
}
