/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAnalyticsPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_insights": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NO_INSIGHTS"},
			"client_insights_sampling": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceClientInsightsSamplingSchema()},
			"client_log_filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceClientLogFilterSchema()},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"full_client_logs": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceFullClientLogsSchema()},
			"metrics_realtime_update": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceMetricsRealTimeUpdateSchema()},
			"significant_log_throttle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"udf_log_throttle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
		},
	}
}
