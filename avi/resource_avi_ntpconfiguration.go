/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceNTPConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ntp_authentication_keys": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceNTPAuthenticationKeySchema()},
			"ntp_server_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema()},
			"ntp_servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceNTPServerSchema()},
		},
	}
}
