package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	StanClusterId = "test-cluster"
	StanClientId  = "1"
	StanSubject   = "foo"
)

var DataDir = "sender/data1/"

func main() {
	if len(os.Args) > 1 {
		DataDir = os.Args[1]
	}
	
	sc, err := stan.Connect(StanClusterId, StanClientId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func() {
		if err := sc.Close(); err != nil {
			log.Error("stan connect close: ", err.Error())
		}
	}()

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
