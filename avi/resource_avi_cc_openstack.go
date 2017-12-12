/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCC_OpenStackSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aapair_ext": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"aapair_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"access_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"admin_access": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"auth_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cfg": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceOpenStackConfigurationSchema()},
			"cfgdrv_ext": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"contrail": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"create_vip_port": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"disc_aasg": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"disc_cfgdrv": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"disc_metasvc": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"disk_sz": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gc_avi": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCC_CronSchema()},
			"gc_cc": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCC_CronSchema()},
			"glance_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ksmgmturl_state": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"metasvc_ext": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mgmt_nw": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCC_MgmtNwSchema()},
			"mgmt_nw_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsxv": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"nuage": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"nvp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"opflex": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"os_role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_role_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_user_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ovs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"park_intfs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCC_ParkIntfSchema()},
			"patch_state": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"portsec_ext": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"rbac_nw_ext": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"schint_ext": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"secgrp_ext": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"svrgrp_ext": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"tenant_access_err": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_aapair_ext": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_anti_affinity": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_cfgdrv": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_contrail": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_extnet": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_intf_secips": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_metasvc": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_nameowner": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_nsxv": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_nuage": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_nuagevip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_nvp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_portsec_ext": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_rbac_nw": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_schint_ext": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_secgrp_ext": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_svrgrp_ext": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}
