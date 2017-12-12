/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceGslbSiteCfgSyncInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"errored_objects": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVersionInfoSchema()},
			"last_changed_time": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceTimeStampSchema()},
			"sync_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
