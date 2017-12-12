/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDebugServiceEngineSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu_shares": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDebugSeCpuSharesSchema()},
			"flags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDebugSeDataplaneSchema()},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "VM name unknown"},
			"seagent_debug": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDebugSeAgentSchema()},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
		},
	}
}
