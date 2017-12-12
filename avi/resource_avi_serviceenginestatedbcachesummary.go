/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceServiceEngineStateDBCacheSummarySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"last_updated_time": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"summary": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceServiceEngineRuntimeSummarySchema()},
		},
	}
}
