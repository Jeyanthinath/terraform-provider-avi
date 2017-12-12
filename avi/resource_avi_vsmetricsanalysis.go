/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVsMetricsAnalysisSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"notes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"se_metrics_issues": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsSeMetricsAnalysisSchema()},
		},
	}
}
