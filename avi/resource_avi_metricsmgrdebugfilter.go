/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricsMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"disable_hw_training": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_grace_period": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_first_n": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"logging_freq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"skip_cluster_map_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"skip_metrics_db_writes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
