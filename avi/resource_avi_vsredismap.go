/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVsRedisMapSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"redis_db": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"redis_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"redis_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true},
			"vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
