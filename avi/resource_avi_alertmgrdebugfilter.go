/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAlertMgrDebugFilterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alert_objid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alert_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cfg_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
