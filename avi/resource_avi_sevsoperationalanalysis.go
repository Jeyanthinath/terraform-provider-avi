/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeVsOperationalAnalysisSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"notes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"pool": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePoolOperationalAnalysisSchema()},
			"server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServerOperationalAnalysisSchema()},
			"vs_oper_status": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceOperationalStatusSchema()},
			"vs_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
