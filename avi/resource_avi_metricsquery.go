/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricsQuerySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aggregate_entity": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"aggregate_obj_id": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"asn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"attack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"browser": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_insights": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ACTIVE"},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"detailed_header": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"devtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dimension_aggregation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_DIMENSION_AGG_NONE"},
			"dimension_filter_op": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_FILTER_EQUALS"},
			"dimension_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"dimension_sampling": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dimensions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"entity_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"filters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricsFiltersSchema()},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
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
			"ipgroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"lang": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"metric_entity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"metric_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"microservice_levels": &schema.Schema{
				Type:     schema.TypeInt,
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
			"os": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pad_missing_data": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
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
			"prediction": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"result_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_FORMAT_JSON"},
			"sampling_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_SAMPLING_FORCE"},
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
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
			"validate_data": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
