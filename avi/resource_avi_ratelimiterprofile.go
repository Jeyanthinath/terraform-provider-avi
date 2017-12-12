/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceRateLimiterProfileSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_ip_connections_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRateProfileSchema()},
			"client_ip_failed_requests_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRateProfileSchema()},
			"client_ip_requests_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRateProfileSchema()},
			"client_ip_scanners_requests_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRateProfileSchema()},
			"client_ip_to_uri_failed_requests_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRateProfileSchema()},
			"client_ip_to_uri_requests_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRateProfileSchema()},
			"http_header_rate_limits": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceRateProfileSchema()},
			"uri_failed_requests_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRateProfileSchema()},
			"uri_requests_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRateProfileSchema()},
			"uri_scanners_requests_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceRateProfileSchema()},
		},
	}
}
