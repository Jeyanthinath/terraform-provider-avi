/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVirtualServiceRelatedEntitiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pools": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricsEntityDataSchema()},
			"servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricsEntityDataSchema()},
			"serviceengines": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMetricsEntityDataSchema()},
			"virtualservice_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
