/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceClientLogConfigurationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_significant_log_collection": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"filtered_log_processing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND"},
			"non_significant_log_processing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND"},
			"significant_log_processing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LOGS_PROCESSING_SYNC_AND_INDEX_ON_DEMAND"},
		},
	}
}
