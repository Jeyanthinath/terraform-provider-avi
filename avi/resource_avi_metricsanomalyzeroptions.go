/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricsAnomalyzerOptionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"anomaly_interpretation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "COMBINE_MODELS_ALL"},
			"anomaly_model_opts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricsAnomalyModelOptionsSchema()},
			"skip_eval_period": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
