/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAuthMappingRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"assign_role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"assign_tenant": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_match": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAuthMatchAttributeSchema()},
			"group_match": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceAuthMatchGroupMembershipSchema()},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"is_superuser": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"role_attribute_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"role_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"tenant_attribute_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
		},
	}
}
