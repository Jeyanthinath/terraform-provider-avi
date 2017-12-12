/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceLogQuerySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adf": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"adf_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cols": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"debug": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"download": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"end": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"expstep": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "json"},
			"groupby": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_rsynced_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nf": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"orderby": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"page": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"page_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"qrtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "QUERY_RANGE_TIME"},
			"query_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"start": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"step": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0"},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "QUERY_CONN_LOGS"},
			"udf": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"virtualservice": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
