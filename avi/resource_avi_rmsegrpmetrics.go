/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRmSeGrpMetricsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"con_metrics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRmConMetricsSchema()},
			"res_metrics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRmResMetricsSchema()},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
		},
	}
}
