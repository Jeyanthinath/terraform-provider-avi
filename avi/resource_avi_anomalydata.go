/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAnomalyDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"anomaly_context": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAnomalyContextSchema()},
			"anomaly_model": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_id_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_OBJ_ID_TYPE_SERVER"},
			"pool_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"value": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true},
		},
	}
}
