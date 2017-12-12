/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricsDerivationDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"derivation_fn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"exclude_derived_metric": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"include_derivation_metrics": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"join_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric_ids": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"result_has_additional_fields": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"skip_backend_derivation": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
