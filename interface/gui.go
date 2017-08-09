package main

import (
	"fmt"
	"github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/init"
	"github.com/solomondg/golang-cube/rpc"
	// "image/color"
	//	"reflect"
	"runtime"
	"sync"
	// "time"
)

const (
	xWinSize = 512
	yWinSize = 512
)

func main() {
	var drawCommand *rpc.GuiDrawCommand // This will be the drawing command that we'll refer back to. The RPC command will simply update this.
	var drawCommandMutex *sync.Mutex    // Prevents two things accessing draw command simultaniously
	var wdeWindow wde.Window            // Main window
	var eventChannel <-chan interface{} // The event channel (so we can access outside of windowSetup)
	var err error                       // So we can use the global stuff with _, err returns

	fmt.Println("Starting GUI + server")

	// Setup window
	windowSetup(&drawCommand, &drawCommandMutex, &wdeWindow, &eventChannel, err)

	stopLoops := false
	go eventLoop(&stopLoops, &wdeWindow)

	wde.Run()
}

func eventLoop(done *bool, wdeWin *wde.Window) {
	events := (*wdeWin).EventChan()
loop: // Label, break loop goes here then skips the next loop
	for event := range events {
		runtime.Gosched()
		switch ev := event.(type) {
		case wde.KeyDownEvent:
			broadcastKeyDown(ev.Key)
		case wde.KeyUpEvent:
			broadcastKeyUp(ev.Key)
		case wde.CloseEvent:
			endProgram(done)
			break loop // Breaks out of whole loop
		}
	}
}

func drawLoop() {

}

func windowSetup(drawComm **rpc.GuiDrawCommand, drawCommMutx **sync.Mutex, wdeWin *wde.Window, evChan *<-chan interface{}, err error) {
	*wdeWin, err = wde.NewWindow(xWinSize, yWinSize)
	handleError(err)

	(*wdeWin).SetTitle("3D")
	(*wdeWin).SetSize(xWinSize, yWinSize)
	(*wdeWin).Show()

	*evChan = (*wdeWin).EventChan()
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
	return
}

func broadcastKeyDown(key string) {
	fmt.Println("Got key down ", key)
}

func broadcastKeyUp(key string) {
	fmt.Println("Got key up ", key)
}

func endProgram(globalEnd *bool) {
	fmt.Println("Ending program")
	*globalEnd = true
}
