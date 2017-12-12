/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMostDifficultObjectEverSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bool_data": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true},
			"enum_data": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"int_data": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"nonref_2": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceNonReferencable2Schema()},
			"optional_nonref_2": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceNonReferencable2Schema()},
			"optional_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ref_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"repeated_nonref_2": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceNonReferencable2Schema()},
			"repeated_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"string_data": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
