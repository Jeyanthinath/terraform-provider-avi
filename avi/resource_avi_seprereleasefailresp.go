/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSePreReleaseFailRespSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"migrate_primary": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Failed to release SE"},
			"se_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"syserr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SYSERR_FAILURE"},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
