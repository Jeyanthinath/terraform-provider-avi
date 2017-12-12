/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceLinuxServerConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceLinuxServerHostSchema()},
			"se_inband_mgmt": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_log_disk_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_log_disk_size_gb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"se_sys_disk_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"se_sys_disk_size_gb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"ssh_user_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
