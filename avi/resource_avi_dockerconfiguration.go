/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceDockerConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_sync_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"ca_tls_key_and_certificate_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_tls_key_and_certificate_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"container_port_match_http_service": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"coredump_directory": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/var/lib/systemd/coredump"},
			"disable_auto_backend_service_sync": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_frontend_service_sync": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_auto_se_creation": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"docker_registry_se": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceDockerRegistrySchema()},
			"east_west_placement_subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrPrefixSchema()},
			"enable_event_subscription": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"feproxy_container_port_as_service": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"feproxy_vips_enable_proxy_arp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"fleet_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_container_ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt}},
			"se_deployment_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_CREATE_SSH"},
			"se_exclude_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema()},
			"se_include_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema()},
			"se_spawn_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  25,
			},
			"se_volume": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/opt/avi"},
			"services_accessible_all_interfaces": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssh_user_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ucp_nodes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"use_container_ip_port": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_controller_image": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
