/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceExponentialMovingAverageCtxSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"average": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true},
			"deviation": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true},
			"max_std_dev": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true},
			"normal_high": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"normal_low": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"prediction": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true},
			"prediction_interval_high": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"prediction_interval_low": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"std_dev": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true},
		},
	}
}
