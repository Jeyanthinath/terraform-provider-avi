/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceTcpcbInfoSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cc_info": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceCcInfoSchema()},
			"irs": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"iss": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_ack_sent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_snd_window": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"np_auto_enabled": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rcv_adv": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rcv_numsacks": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rcv_nxt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rcv_scale": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rcv_urg_pointer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rcv_window": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"request_r_scale": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rfbuf_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rfbuf_ts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sack_newdata": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sackblks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSackBlkSchema()},
			"sackhint": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      func(v interface{}) int { return 0 }, Elem: ResourceSackHintSchema()},
			"snd_conj_window": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snd_cwnd_prev": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snd_fack": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snd_limited": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snd_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snd_numholes": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snd_nxt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snd_recover": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snd_recover_prev": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snd_scale": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snd_ssthresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snd_ssthresh_prev": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snd_una": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snd_urg_pointer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snd_window": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snd_wl1": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snd_wl2": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_2msl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_badrxtwin": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_bytes_acked": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_conn_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_dupacks": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_finwait2": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_flags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_iobc": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_keepcnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_keepidle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_keepinit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_keepintvl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_max_rxtshift": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_max_syn_rxtshift": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_maxidle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_maxopd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_maxseg": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_oobflags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_quickacks": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_rcvoopack": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_rcvtime_tick": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_rcvzerowin": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_rtseq": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_rtt_avi": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_rtt_variance": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_rttbest": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_rttlow": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_rttmin": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_rtttime_tick": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_rttupdated": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_rxtcur_tick": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_rxtshift": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_segqlen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_sendprobe": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"t_smoothed_rtt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_sndrexmitpack": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_sndtimeorexmit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_sndzerowin": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_softerror": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_starttime_tick": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_timewait": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"t_usr_flags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_state": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tcp_state_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ts_offset": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ts_recent": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ts_recent_age": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
