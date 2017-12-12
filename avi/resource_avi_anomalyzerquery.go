/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAnomalyzerQuerySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aggregate_obj_id": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"aggregation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_ANOMALY_AGG_NONE"},
			"aggregation_window": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"detailed_header": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"dimension_filter_op": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_FILTER_EQUALS"},
			"entity_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"include_refs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"include_related": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  50,
			},
			"metric_entity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"metric_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"model": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"order_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "metric_timestamp"},
			"page": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pool_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ANZ_PRIORITY_HIGH"},
			"result_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_FORMAT_JSON"},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"serviceengine_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "1970-01-01T00:00:00"},
			"step": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"stop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
