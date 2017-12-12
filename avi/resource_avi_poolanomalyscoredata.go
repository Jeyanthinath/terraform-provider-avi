/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcePoolAnomalyScoreDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"anomalous_l4_metrics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"anomalous_l7_metrics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"avg_server_anomaly_score": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"per_anomalous_l4_metrics": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true},
			"per_anomalous_l7_metrics": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"reason_attr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_anomaly_scores": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServerAnomalyScoreSchema()},
		},
	}
}
