/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricsOptionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"archive_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_DB_TABLE_ARCHIVE_ALL"},
			"entity_types": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"indexes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"metrics_aggregation": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"metrics_obj_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metrics_object": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"metrics_z_score_threshold": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"table_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
