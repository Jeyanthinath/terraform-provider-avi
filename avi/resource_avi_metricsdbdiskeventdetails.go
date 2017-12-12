/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMetricsDbDiskEventDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metrics_deleted_tables": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"metrics_free_sz": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"metrics_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
		},
	}
}
