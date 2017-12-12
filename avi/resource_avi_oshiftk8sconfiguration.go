/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceOShiftK8SConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_sync_frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"avi_bridge_subnet": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrPrefixSchema()},
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
			"default_service_as_east_west_service": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
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
			"disable_auto_gs_sync": &schema.Schema{
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
			"l4_health_monitoring": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"master_nodes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"node_availability_zone_label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ns_exclude_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema()},
			"ns_include_attributes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceMesosAttributeSchema()},
			"nuage_controller": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceNuageSDNControllerSchema()},
			"sdn_overlay": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
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
			"secure_egress_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"service_account_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_port_match_http_service": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"shared_virtualservice_namespace": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ssh_user_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_controller_image": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_scheduling_disabled_nodes": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_service_cluster_ip_as_ew_vip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
