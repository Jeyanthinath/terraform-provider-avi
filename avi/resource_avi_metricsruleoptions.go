/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricsRuleOptionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"event_id_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"high_wms": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeFloat}},
			"low_wms": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeFloat}},
			"only_positive_anomalies": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"rule_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"valid_event_period": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
