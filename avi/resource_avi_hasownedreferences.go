/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHasOwnedReferencesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"multiple_ref": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIsOwnedReferenceSchema()},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
		},
	}
}
