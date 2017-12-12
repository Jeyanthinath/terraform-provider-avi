/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVsCandidateSeHostListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"can_spawn_new_se": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsCandidateHostSchema()},
			"se": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVsCandidateSeSchema()},
		},
	}
}
