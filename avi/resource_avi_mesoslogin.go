/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceMesosLoginSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"include_mesos_hosts": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"marathon_logins": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMarathonLoginSchema()},
			"mesos_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
		},
	}
}
