/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceLdapAuthSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"base_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bind_as_administrator": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"email_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "email"},
			"full_name_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "name"},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  389,
			},
			"security_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"settings": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceLdapDirectorySettingsSchema()},
			"user_bind": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceLdapUserBindSettingsSchema()},
		},
	}
}
