/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHardwareSecurityModuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"nethsm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHSMThalesNetHsmSchema()},
			"rfs": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceHSMThalesRFSSchema()},
			"sluna": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceHSMSafenetLunaSchema()},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
