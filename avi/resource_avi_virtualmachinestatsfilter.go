/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVirtualMachineStatsFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cpu_utilization": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"some_other_metric": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
