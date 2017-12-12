/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAnomalyzerQueryResponseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metric_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"series": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAnomalyDataSeriesSchema()},
			"start": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"step": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"stop": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
