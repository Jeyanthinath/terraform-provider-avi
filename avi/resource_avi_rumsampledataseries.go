/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRumSampleDataSeriesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRumSampleDataSchema()},
			"header": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceRumSampleDataHeaderSchema()},
		},
	}
}
