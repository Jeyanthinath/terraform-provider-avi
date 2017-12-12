/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAlertRuleMetricSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"metric_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_threshold": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceAlertMetricThresholdSchema()},
		},
	}
}
