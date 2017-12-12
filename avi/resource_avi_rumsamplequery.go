/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRumSampleQuerySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"browser": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"devtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dimension_filter_op": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_FILTER_EQUALS"},
			"dimensions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"entity_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"include_refs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
				Default:  5,
			},
			"metric_entity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"os": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prediction": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"result_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "METRICS_FORMAT_JSON"},
			"sample_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "1970-01-01T00:00:00"},
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
		},
	}
}
