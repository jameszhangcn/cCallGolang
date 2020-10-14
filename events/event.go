package event

/*
#cgo LDFLAGS:
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include "event.h"

extern int mdp_event_end_marker_handler(uint32_t ueID, uint8_t rbID, uint8_t tnl_type);
extern int mdp_event_gtp_err_handler(uint32_t ueID, uint32_t teID, uint8_t rbID);

void mdp_event_init(){
	//TODO add mdp lib
    // mdp_event_slave_init();
    //mdp_event_slave_setup(NULL, (mdp_event_recv_t*)mdp_event_handler, false);
	printf("\r\n mdp_event_init \r\n");
	return;
}

int mdp_event_handler(void *data, uint32_t len, void* private){
	mdp_event_hdr_t* eventHdr = (mdp_event_hdr_t*)data;
	printf("\r\n Event type %d, len %d, thread %d", eventHdr->type, eventHdr->len, eventHdr->thread_index);

	cu_up_event_t* upEvent = (cu_up_event_t*)((uint8_t*)data+sizeof(mdp_event_hdr_t));
    printf("\r\n CUUP Event type %d ", upEvent->event);

	switch(upEvent->event)
	{
		case CU_UP_EVENT_END_MARKER_IND:
		{
			printf("CU_UP_EVENT_END_MARKER_IND");
			cu_up_end_marker_ind_t * end_marker = (cu_up_end_marker_ind_t*)malloc(sizeof(cu_up_end_marker_ind_t));
			*end_marker = upEvent->end_marker_ind;
            mdp_event_end_marker_handler(end_marker->ue_id, end_marker->rb_id, end_marker->tnl_type);
			break;
		}
		case CU_UP_EVENT_GTP_ERR_IND:
		{
			printf("CU_UP_EVENT_GTP_ERR_IND");
			cu_up_gtp_err_ind_t * gtp_err = (cu_up_gtp_err_ind_t*)malloc(sizeof(cu_up_gtp_err_ind_t));
			*gtp_err = upEvent->gtp_err_ind;
            mdp_event_gtp_err_handler(gtp_err->ue_id, gtp_err->te_id, gtp_err->rb_id);

			break;
		}
		default:
			printf("Unsupported event %d", upEvent->event);
			break;
	}

	return 0;
}

void mdp_event_start_periodic(){
    char data[100];
    char private[20];
    int len = 20;
	printf("\r\n mdp_event_start_periodic \r\n ");
	cu_up_event_t *event = (cu_up_event_t*)malloc(sizeof(cu_up_event_t)+sizeof(cu_up_end_marker_ind_t));
	event->event = CU_UP_EVENT_END_MARKER_IND;
	cu_up_end_marker_ind_t ind;
	ind.ue_id = 1234;
	ind.rb_id = 1;
	ind.tnl_type = 1;
	memcpy(&(event->end_marker_ind),&ind,sizeof(cu_up_end_marker_ind_t));
	memcpy(data+sizeof(mdp_event_hdr_t),event,sizeof(cu_up_event_t)+sizeof(cu_up_end_marker_ind_t));
	mdp_event_handler((void*)data,len,(void*)private);

	cu_up_event_t *event1 = (cu_up_event_t*)malloc(sizeof(cu_up_event_t)+sizeof(cu_up_gtp_err_ind_t));
	event1->event = CU_UP_EVENT_GTP_ERR_IND;
	cu_up_gtp_err_ind_t ind1;
	ind1.ue_id = 4321;
	ind1.rb_id = 1;
	ind1.te_id = 2345;
	memcpy(&(event1->gtp_err_ind),&ind1,sizeof(cu_up_gtp_err_ind_t));
	memcpy(data+sizeof(mdp_event_hdr_t),event1,sizeof(cu_up_event_t)+sizeof(cu_up_gtp_err_ind_t));
	mdp_event_handler((void*)data,len,(void*)private);

	return;
}
*/
import "C"
import (
	"time"
)

const MAX_EVENT_DATA_LEN = 1024

var eventData [MAX_EVENT_DATA_LEN]byte
var eventPrivate [MAX_EVENT_DATA_LEN]byte

func MdpEventInit() {

	C.mdp_event_init()

	//for testing
	go func() {
		for {
			time.Sleep(time.Second)
			C.mdp_event_start_periodic()
		}
	}()
}