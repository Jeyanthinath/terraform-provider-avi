/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"log"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceSeConsumerProtoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active_standby_se_tag": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ACTIVE_STANDBY_SE_1"},
		"add_vnic_se_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"apic_mode": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"at_spawn_limit": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"attach_ip_fail_reason": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"attach_ip_fail_reason_code": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SYSERR_FAILURE"},
		"availability_zone": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"bootup_se_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"cloud_apis_ready": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"cloud_apis_unready_reason": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"cloud_apis_unready_reason_code": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SYSERR_FAILURE"},
		"cloud_ready": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"cloud_unready_reason": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"cloud_unready_reason_code": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SYSERR_FAILURE"},
		"cookie": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"create_disabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"disabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"dns_info": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsInfoSchema()},
		"east_west_placement": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"enable_rhi": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"enable_rhi_snat": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"exclude_create_op": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"exclude_vnic_op": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"fip": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrSchema()},
		"ign_pool_net_reach": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"insufficient_buffer_se": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"lifs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString}},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"next_spawn_ticks": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"notify_on_avail": &schema.Schema{
			Type:     schema.TypeBool,
			Required: true},
		"notify_q_name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true},
		"num_se": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true},
		"num_se_pending": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true},
		"num_seq_attach_ip_fail": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"num_seq_spawn_fail": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"num_seq_vnic_fail": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"num_spawn_attempts": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"num_stby_se": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true},
		"num_stby_se_pending": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true},
		"num_vnics": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"op_start_ticks": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"parent_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"pending_vnic_op": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"pending_vnic_op_reason": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"placement_disabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"placement_disabled_reason": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"placement_disabled_reason_code": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SYSERR_FAILURE"},
		"placement_history": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePlacementHistorySchema()},
		"query_host_pending": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"res_pending": &schema.Schema{
			Type:     schema.TypeSet,
			Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceVsResProfileSchema()},
		"res_total": &schema.Schema{
			Type:     schema.TypeSet,
			Required: true, Set: func(v interface{}) int { return 0 }, Elem: ResourceVsResProfileSchema()},
		"resources_consumed": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceSeResourceConsumedProtoSchema()},
		"scaleout_ecmp": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"se_group_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"se_placement_pending": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceSeResourceConsumedProtoSchema()},
		"servers": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceConServerSchema()},
		"service_subnet": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceIpAddrPrefixSchema()},
		"services": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceVsProtocolSchema()},
		"snat_ip": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrSchema()},
		"spawn_fail_reason": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"spawn_fail_reason_code": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SYSERR_FAILURE"},
		"static_se_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"static_se_uuids": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString}},
		"svc_obj_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"svc_q_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"tenant_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"use_res_profile": &schema.Schema{
			Type:     schema.TypeBool,
			Required: true},
		"vip": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceConVipSchema()},
		"vnic_ip_disabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"vnic_op_disabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"vrf_context_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"write_ops_disabled_info": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceVipPlacementResolutionInfoSchema()},
		"write_ops_disabled_reason": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"write_ops_disabled_reason_code": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SYSERR_FAILURE"},
	}
}

func resourceAviSeConsumerProto() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSeConsumerProtoCreate,
		Read:   ResourceAviSeConsumerProtoRead,
		Update: resourceAviSeConsumerProtoUpdate,
		Delete: resourceAviSeConsumerProtoDelete,
		Schema: ResourceSeConsumerProtoSchema(),
	}
}

func ResourceAviSeConsumerProtoRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSeConsumerProtoSchema()
	log.Printf("[INFO] ResourceAviSeConsumerProtoRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/seconsumerproto/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	// no need to set the ID
	log.Printf("ResourceAviSeConsumerProtoRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviSeConsumerProtoRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviSeConsumerProtoRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviSeConsumerProtoRead Updated %v\n", d)
	return nil
}

func resourceAviSeConsumerProtoCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSeConsumerProtoSchema()
	err := ApiCreateOrUpdate(d, meta, "seconsumerproto", s)
	log.Printf("[DEBUG] created object %v: %v", "seconsumerproto", d)
	if err == nil {
		err = ResourceAviSeConsumerProtoRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "seconsumerproto", d)
	return err
}

func resourceAviSeConsumerProtoUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSeConsumerProtoSchema()
	err := ApiCreateOrUpdate(d, meta, "seconsumerproto", s)
	if err == nil {
		err = ResourceAviSeConsumerProtoRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "seconsumerproto", d)
	return err
}

func resourceAviSeConsumerProtoDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "seconsumerproto"
	log.Println("[INFO] ResourceAviSeConsumerProtoRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil {
			log.Println("[INFO] resourceAviSeConsumerProtoDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
