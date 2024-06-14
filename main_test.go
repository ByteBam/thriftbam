package main

import (
	"flag"
	"github.com/ByteBam/thirftbam/pkg/utils/config"
	"testing"
)

func TestName(t *testing.T) {
	var envConf = flag.String("conf", "biz/config/config.yaml", "config file path")
	flag.Parse()
	conf := config.NewConfig(*envConf)
	c := conf.Get("rocketmq.consumers").([]interface{})
	for _, v := range c {
		t.Log(v.(string))
	}
	t.Log(c)
}
