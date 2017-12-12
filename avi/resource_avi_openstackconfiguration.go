/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceOpenStackConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_tenant": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"admin_tenant_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allowed_address_pairs": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"anti_affinity": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"auth_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"config_drive": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"contrail_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"contrail_plugin": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"external_networks": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"free_floatingips": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"hypervisor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "KVM"},
			"hypervisor_properties": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceOpenStackHypervisorPropertiesSchema()},
			"img_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "OS_IMG_FMT_AUTO"},
			"import_keystone_tenants": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"insecure": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"intf_sec_ips": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"keystone_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"map_admin_to_cloudadmin": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"mgmt_network_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"mgmt_network_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_owner": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"neutron_rbac": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"nuage_organization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  8443,
			},
			"nuage_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nuage_virtualip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"nuage_vsd_host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_security": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"privilege": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"prov_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"role_mapping": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceOpenStackRoleMappingSchema()},
			"se_group_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_groups": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"tenant_se": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"usable_network_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"use_admin_url": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"use_internal_endpoints": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_keystone_auth": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"use_nuagevip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
