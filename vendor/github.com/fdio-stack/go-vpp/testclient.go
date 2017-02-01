package govpp

/*
#cgo CFLAGS: -I/usr/local/include/libvpp_cgoclient
#cgo LDFLAGS: -lvpp_cgoclient
#include <vpp_client.h>
*/
import "C"
import (
	"./srv"
	"github.com/briandowns/spinner"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func myspinner() {
	s := spinner.New(spinner.CharSets[34], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                    // Start the spinner
	time.Sleep(5 * time.Second)                                  // Run for some time to simulate work
	s.Stop()
}

/***************** MAIN ******************/
func main() {

	/* This block loops until Ctrl-C is hit then disconnects */
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		govppp.VppDisconnect()
		os.Exit(1)
	}()
	/* END clean up on SIGINT */
	govppp.VppConnect("vpp_contiv_client")
	govppp.VppAddInterface("web1")
	govppp.VppAddInterface("web2")
	govppp.VppAddInterfaceIp("web1", "192.199.1.1/24")
	govppp.VppAddInterfaceIp("web2", "192.199.2.1/24")
	for {
		myspinner()
	}
}
