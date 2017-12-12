/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcePGDeploymentRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metric_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "health.health_score_value"},
			"operator": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CO_GE"},
			"threshold": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
