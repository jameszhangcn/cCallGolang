package event

import (
	"C"
	"fmt"
)

var (
	eventsCounter     = 0
	endMarkerMethodId int32
	gtpErrorMethodId  int32
)

type stEndMarkerInd struct {
	ue_id    uint32
	rb_id    uint8
	tnl_type uint8
}

type stGtpErrInd struct {
	ue_id uint32
	te_id uint32
	rb_id uint8
}

//export mdp_event_end_marker_handler
func mdp_event_end_marker_handler(ueID C.ulong, rbID C.uchar, tnlType C.uchar) {
	fmt.Println("Received event end marker, ueID ", ueID)
	ind := &stEndMarkerInd{
		ue_id:    uint32(ueID),
		rb_id:    uint8(rbID),
		tnl_type: uint8(tnlType),
	}
	fmt.Println(ind)
}

//export mdp_event_gtp_err_handler
func mdp_event_gtp_err_handler(ueID C.ulong, teID C.ulong, rbID C.uchar) {
	fmt.Println("Received event gtp err, ueID", ueID)
	ind := &stGtpErrInd{
		ue_id: uint32(ueID),
		te_id: uint32(teID),
		rb_id: uint8(rbID),
	}
	fmt.Println(ind)
}