package main

import (
	"flag"
	"fmt"
	"testing"
)

var systemTest *bool
var configDirPath string

func init() {
	systemTest = flag.Bool("systemTest", false, "Set to true when running system tests")
	flag.StringVar(&configDirPath, "configDir", "", "path to additional filebeat configuration directory with .yml files")
}

//Test started when the test binary is started. Only calls main.
func TestSystem(t *testing.T) {
	if *systemTest {
		main()
		fmt.Println("returned from main")
	}
}
