/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceruleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "reject"},
			"appliedtolist": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcensxAppliedToListSchema()},
			"destinations": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcensxRuleDestsSchema()},
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"logged": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"packettype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sectionid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"services": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcensxServicesSchema()},
			"sources": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcensxRuleSrcsSchema()},
		},
	}
}
