/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceAdminAuthConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow_local_user_login": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"auth_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mapping_rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceAuthMappingRuleSchema()},
		},
	}
}
