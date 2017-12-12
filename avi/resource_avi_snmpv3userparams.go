/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSnmpV3UserParamsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_passphrase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "avinetworks"},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SNMP_V3_AUTH_MD5"},
			"priv_passphrase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "avinetworks"},
			"priv_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SNMP_V3_PRIV_DES"},
		},
	}
}
