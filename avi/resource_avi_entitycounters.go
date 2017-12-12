/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceEntityCountersSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"counters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVcenterPerfCounterSchema()},
			"disk1_usage": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"disk2_usage": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"disk3_usage": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"disk4_usage": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"interested": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVmwarePerfMetricIdSchema()},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"refreshrate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"seq_no": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"supported": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVmwarePerfMetricIdSchema()},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
		},
	}
}
