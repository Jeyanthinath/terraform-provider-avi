/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHealthScoreSecurityDataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dos_attack_level_data": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceDosAttackLevelDataSchema()},
			"reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "No Security Penalty"},
			"reason_attr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_dos_attack_level_data": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceServiceEngineDosAttackLevelDataSchema()},
			"serviceengine_dos_attack_level_data": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceServiceEngineDosAttackLevelDataSchema()},
			"ssl_score_data": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSslScoreDataSchema()},
			"value": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
		},
	}
}
