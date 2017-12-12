/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHttpCacheETypeStatsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"count": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"encoding": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
