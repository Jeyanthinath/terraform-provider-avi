/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHTTPSecurityRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceHTTPSecurityActionSchema()},
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"log": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceMatchTargetSchema()},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
