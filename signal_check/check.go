package signal_check

import (
    "os"
    "os/signal"
)

var signalChannel chan os.Signal
var callback func()

func Init(){
    signalChannel = make(chan os.Signal)
    signal.Notify(signalChannel)
}

func SetSignalCallBack(callbackProc func()){
    callback = callbackProc
}

func Wait(){
    <- signalChannel
    if nil != callback {callback()}
}