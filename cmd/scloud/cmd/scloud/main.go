package main

//go:generate go run gen_version.go

import (
	"github.com/Laplace-Transformer/splunk-cloud-sdk-go/cmd/scloud/cmd"
	"github.com/golang/glog"
)

func main() {
	cmd.Execute()
	glog.Flush()
}
