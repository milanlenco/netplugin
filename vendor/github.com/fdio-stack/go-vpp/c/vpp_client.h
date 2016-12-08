#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/mman.h>
#include <sys/stat.h>
#include <netinet/in.h>
#include <signal.h>
#include <pthread.h>
#include <unistd.h>
#include <time.h>
#include <fcntl.h>
#include <string.h>
#include <arpa/inet.h>

#include <vppinfra/clib.h>
#include <vppinfra/vec.h>
#include <vppinfra/hash.h>
#include <vppinfra/bitmap.h>
#include <vppinfra/fifo.h>
#include <vppinfra/time.h>
#include <vppinfra/mheap.h>
#include <vppinfra/heap.h>
#include <vppinfra/pool.h>
#include <vppinfra/format.h>
#include <vppinfra/error.h>

#include <vnet/vnet.h>
#include <vlib/vlib.h>
#include <vlib/unix/unix.h>
#include <vlibapi/api.h>
#include <vlibmemory/api.h>

#include <vpp-api/vpe_msg_enum.h>

#include <vnet/ip/ip.h>
#include <vnet/interface.h>

#define vl_typedefs             /* define message structures */
#include <vpp-api/vpe_all_api_h.h>
#undef vl_typedefs

typedef struct {
    int link_events_on;
    int stats_on;
    int oam_events_on;
    /* convenience */
    unix_shared_memory_queue_t * vl_input_queue;
    u32 my_client_index;
    char *my_client_name;
} client_main_t;

typedef struct vpp_interface_counters_record {
    struct timespec timestamp;
    int sw_if_index;
    char* counter_name;
    u64 counter;
    struct vpp_interface_counters_record *next;
} vpp_interface_counters_record_t;

typedef struct vpp_interface_summary_counters_record {
    struct timespec timestamp;
    int sw_if_index;
    char* counter_name;
    u64 packet_counter;
    u64 byte_counter;
    struct vpp_interface_summary_counters_record *next;
} vpp_interface_summary_counters_record_t;


client_main_t cm;

void set_flags (int *sw_if_index, int *up_down, client_main_t *cm);
void l2_patch_add_del (client_main_t *tm, int is_add);
void get_vpp_summary_stats(client_main_t *cm);
void add_l2_bridge (int bd_id, client_main_t *cm);
void set_l2_bridge_interface (int bd_id, int *rx_if_index, client_main_t *cm);
void add_af_packet_interface (char *intf, client_main_t *cm);
void add_del_interface_address (int enable_disable, int *sw_if_index, u32 *ipaddr, u8 *length, client_main_t *cm);
void stats_enable_disable (int enable, client_main_t *cm);

int connect_to_vpp(client_main_t *cm);
int disconnect_from_vpp(void);

/* Callbacks to GO functions - must have //export Gocallback_ above GO func declation */
extern void gocallback_l2_bridge_reply (int *retval);
extern void gocallback_l2_bridge_set_interface_reply (int *retval);
extern void gocallback_af_packet_create_reply (int *retval, int *sw_if_index);
extern void gocallback_add_del_address_reply ();
extern void gocallback_set_interface_flags (int *retval);
extern void gocallback_connect_to_vpp (client_main_t *cm);
extern void gocallback_vnet_summary_interface_counters (int *record_count, vpp_interface_summary_counters_record_t *records);
extern void gocallback_vnet_interface_counters (int *record_count, vpp_interface_counters_record_t *records);
