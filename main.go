package main

import (
	"flag"
	"k8s.io/klog/v2"
	"github.com/shappy0/ntui/cmd"
)

func main() {
	klog.InitFlags(nil)
	if Err := flag.Set("logtostderr", "false"); Err != nil {
		panic(Err)
	}
	if Err := flag.Set("alsologtostderr", "false"); Err != nil {
		panic(Err)
	}
	if Err := flag.Set("stderrthreshold", "fatal"); Err != nil {
		panic(Err)
	}
	if Err := flag.Set("v", "0"); Err != nil {
		panic(Err)
	}
	klog.Info("Ok")
	cmd.Run()
}