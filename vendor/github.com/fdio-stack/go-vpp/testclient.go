package govpp

/*
#cgo CFLAGS: -I/usr/local/include/libvpp_cgoclient
#cgo LDFLAGS: -lvpp_cgoclient
#include <vpp_client.h>
*/
import "C"
import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fdio-stack/go-vpp/srv"
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
		srv.VppDisconnect()
		os.Exit(1)
	}()
	/* END clean up on SIGINT */
	srv.VppConnect("vpp_contiv_client")
	srv.VppAddInterface("web1")
	srv.VppAddInterface("web2")
	srv.VppAddInterfaceIp("web1", "192.199.1.1/24")
	srv.VppAddInterfaceIp("web2", "192.199.2.1/24")
	for {
		myspinner()
	}
}
