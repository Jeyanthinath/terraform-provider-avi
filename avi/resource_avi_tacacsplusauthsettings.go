/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceTacacsPlusAuthSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authorization_attrs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAuthTacacsPlusAttributeValuePairSchema()},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  49,
			},
			"server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AUTH_TACACS_PLUS_SERVICE_LOGIN"},
		},
	}
}
