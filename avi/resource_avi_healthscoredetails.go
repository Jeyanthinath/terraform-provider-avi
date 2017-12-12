/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHealthScoreDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"anomaly_penalty": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"anomaly_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"performance_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"performance_score": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"previous_value": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"resources_penalty": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"resources_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_penalty": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"security_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"step": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sub_resource_prefix": &schema.Schema{
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
