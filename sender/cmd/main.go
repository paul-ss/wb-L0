package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"path/filepath"
)

const (
	StanClusterId = "test-cluster"
	StanClientId  = "1"
	StanSubject   = "foo"

	DataDir = "sender/data/"
)

func main() {
	sc, err := stan.Connect(StanClusterId, StanClientId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer sc.Close()

	fs, err := ioutil.ReadDir(DataDir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, f := range fs {
		data, err := ioutil.ReadFile(filepath.Join(DataDir, f.Name()))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if err = sc.Publish(StanSubject, data); err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
