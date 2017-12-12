/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHealthScoreQueryResponseSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entity_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"metric_entity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"series": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHealthScoreDataSeriesSchema()},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
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
