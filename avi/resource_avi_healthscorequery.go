/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHealthScoreQuerySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"detailed_header": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"dimension_aggregation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_DIMENSION_AGG_NONE"},
			"entity_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"hs_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HEALTH_SCORE"},
			"include_refs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"include_statistics": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"metric_entity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"order_series_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pad_missing_data": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"patch_oper_status": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"pool_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"result_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_FORMAT_JSON"},
			"server": &schema.Schema{
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
				Default:  300,
			},
			"stop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"summary": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"tenant_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
