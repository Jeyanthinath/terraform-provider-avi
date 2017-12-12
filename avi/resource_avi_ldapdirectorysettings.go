/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceLdapDirectorySettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_bind_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "(objectClass=*)"},
			"group_member_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "member"},
			"group_member_is_full_dn": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"group_search_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_search_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AUTH_LDAP_SCOPE_SUBTREE"},
			"ignore_referrals": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"user_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"user_id_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_search_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_search_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AUTH_LDAP_SCOPE_ONE"},
		},
	}
}
