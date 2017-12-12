/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceApplicationResourcesScoreDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"reason_attr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtualservice_resources_scores": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVirtualServiceResourcesScoreSchema()},
		},
	}
}
