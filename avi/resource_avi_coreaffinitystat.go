/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCoreAffinityStatSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"compact_core_nums": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt}},
			"core_nonaffinity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_flow_cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceLbServerSchema()},
			"stretch_factor": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sum_weights": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"weighted_core_s_indices": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRepeatedSrvIdxWeightsSchema()},
		},
	}
}
