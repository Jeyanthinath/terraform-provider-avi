/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeHmEventGslbPoolMemberDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAppInfoSchema()},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"failure_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"shm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSeHmEventShmDetailsSchema()},
			"ssl_error_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
