/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHealthScoreHeaderSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"anomaly_penalty_statistics": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceMetricStatisticsSchema()},
			"contributors": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHealthScoreTypeContributorDataSchema()},
			"entity_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hs_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metrics_min_scale": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"missing_intervals": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricsMissingDataIntervalSchema()},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "health_score"},
			"performance_score_statistics": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceMetricStatisticsSchema()},
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resources_penalty_statistics": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceMetricStatisticsSchema()},
			"security_penalty_statistics": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceMetricStatisticsSchema()},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"serviceengine_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"statistics": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceMetricStatisticsSchema()},
			"units": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRIC_COUNT"},
		},
	}
}
