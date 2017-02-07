package stats

import (
	"encoding/gob"
	"os"
	"time"
)

var ifRecords_filename string = "./interfacestats.txt"
var ifRecords_fh *os.File

type VppInterfaceStats_key_t struct {
	Sw_if_index int
	Timestamp   time.Time
}

type VppInterfaceStats_t struct {
	Key                VppInterfaceStats_key_t
	Packets_tx         int64
	Packets_rx         int64
	Bytes_tx           int64
	Bytes_rx           int64
	Drop               int64
	Punt               int64
	Ip4                int64
	Ip6                int64
	Rx_no_buf          int64
	Rx_miss            int64
	Rx_error           int64
	Tx_error_fifo_full int64
	Bogus              int64
}

func AddInterfaceRecord(record VppInterfaceStats_t) {

	//fmt.Printf("r: %+v \n", record)
	enc := gob.NewEncoder(ifRecords_fh)
	err := enc.Encode(record)
	if err != nil {
		panic(err)
	}
}

func init() {
	var err error
	ifRecords_fh, err = os.Create(ifRecords_filename)
	if err != nil {
		panic(err)
	}
}

func Close() {

	ifRecords_fh.Close()
}
