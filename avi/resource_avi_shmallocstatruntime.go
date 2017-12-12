/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceShMallocStatRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sh_mallocstat_entry": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceShMallocStatEntrySchema()},
		},
	}
}
