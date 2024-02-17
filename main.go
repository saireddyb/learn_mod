package main

import (
	"github.com/saireddyb/learn_mod/pkg/api"
	log "github.com/sirupsen/logrus"
)

func main() {
  log.WithFields(log.Fields{
    "animal": "walrus",
    "number": 1,
    "size":   10,
  }).Info("A walrus appears")
	api.CallAPI()
}