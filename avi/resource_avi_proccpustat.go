/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceProcCpuStatSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"process_cpu_ewma": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"process_cpu_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"process_memory_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"process_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
