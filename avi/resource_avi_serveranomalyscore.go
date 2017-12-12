/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceServerAnomalyScoreSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hs_entity": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceHealthScoreEntitySchema()},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"score_data": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceServerAnomalyScoreDataSchema()},
			"value": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true},
		},
	}
}
