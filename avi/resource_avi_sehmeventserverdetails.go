/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeHmEventServerDetailsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAppInfoSchema()},
			"failure_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
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
