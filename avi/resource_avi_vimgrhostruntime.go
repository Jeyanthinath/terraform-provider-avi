/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceVIMgrHostRuntimeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud"},
			"cluster_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cntlr_accessible": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"connection_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "connected"},
			"cpu_hz": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"maintenance_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"managed_object_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"mem": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mgmt_portgroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"network_uuids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
			"num_cpu_cores": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_cpu_packages": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"num_cpu_threads": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCdpLldpInfoSchema()},
			"powerstate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"quarantine_start_ts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"quarantined": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"quarantined_periods": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"se_fail_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"se_success_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true},
			"vm_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString}},
		},
	}
}
