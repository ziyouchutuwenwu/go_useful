package signal_check

import (
	"fmt"
	"testing"
)


func onShutdown(){
    fmt.Println("on signal")
}

func TestSignal(t *testing.T) {
    Init()
    SetSignalCallBack(onShutdown)
    Wait()
}