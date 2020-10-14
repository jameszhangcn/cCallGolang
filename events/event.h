#ifndef MDPEVENT_H
#define MDPEVENT_H
#include <stdint.h>
typedef enum
{
    CU_UP_EVENT_DL_DATA_IND,
    CU_UP_EVENT_UL_DATA_IND,
    CU_UP_EVENT_END_MARKER_IND,
    CU_UP_EVENT_GTP_ERR_IND,
    CU_UP_EVENT_MAX,
} cu_up_event_e;

typedef struct
{
    uint32_t ue_id;
    uint8_t ppi;
} cu_up_dl_data_ind_t;

typedef struct
{
    uint32_t ue_id;
    uint8_t  rb_id;
    uint8_t  tnl_type;
} cu_up_end_marker_ind_t;

typedef struct
{
    uint32_t ue_id;
    uint32_t te_id;
    uint8_t  rb_id;

} cu_up_gtp_err_ind_t;

typedef struct
{
    cu_up_event_e event;

    union {
        cu_up_dl_data_ind_t dl_data_ind;
        cu_up_gtp_err_ind_t gtp_err_ind;
        cu_up_end_marker_ind_t end_marker_ind;
    };
} cu_up_event_t;

typedef struct 
{
    uint64_t tv_sec;
    uint32_t tv_usec;
    uint32_t thread_index;
    uint8_t type;
    uint16_t len;
}mdp_event_hdr_t;

#endif