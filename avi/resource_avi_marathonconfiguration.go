/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMarathonConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"framework_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"marathon_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"marathon_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "http://leader.mesos:8080"},
			"marathon_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_port_range": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcePortRangeSchema()},
			"public_port_range": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourcePortRangeSchema()},
			"tenant": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "admin"},
			"use_token_auth": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vs_name_tag_framework": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
