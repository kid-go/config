package test

import (
	"fmt"
	"github.com/leor-w/config"
	"github.com/leor-w/config/local"
	"testing"
	"time"
)

type Test struct {
	Hello string
}

func TestLocalConfig(t *testing.T) {
	lProvider := local.New(
		local.WithConfigName("config"),
		local.WithConfigPath("./"),
		local.WithConfigType("yaml"),
	)
	config.New()
	config.SetProvider(lProvider)
	test := &Test{}
	if err := lProvider.ReadConfig(); err != nil {
		t.Error(err.Error())
	}
	if err := lProvider.Unmarshal("test", test); err != nil {
		t.Error(err.Error())
	}
	lProvider.OnWatch(func() {
		if err := lProvider.Unmarshal("test", test); err != nil {
			t.Error(err.Error())
		}
	})
	for {
		time.Sleep(time.Second)
		fmt.Println(test.Hello)
	}
}
