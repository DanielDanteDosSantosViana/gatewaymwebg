package test

import (
	"github.com/DanielDanteDosSantosViana/gatewaymwebg/config"
	"testing"
)

func TestConfigFileMustBeLoaded(t *testing.T) {
	configuration := config.New()
	_, err := configuration.Load()
	if err != nil {
		t.Errorf("Error Config don't working :  %v", err)
	}

}

func TestIPMustBeNonEmpty(t *testing.T) {
	configuration := config.New()
	conf, _ := configuration.Load()
	ip := conf.Services[0].Ip
	if ip == "" {
		t.Errorf("Error ip is Empty :  %s", ip)
	}
}

func TestHostMustBeNonEmpty(t *testing.T) {
	configuration := config.New()
	conf, _ := configuration.Load()
	host := conf.Services[0].Host
	if host == "" {
		t.Errorf("Error host is Empty :  %s", host)
	}
}

func TestMethodMustBeNonEmpty(t *testing.T) {
	configuration := config.New()
	conf, _ := configuration.Load()
	method := conf.Services[0].Method
	if method == "" {
		t.Errorf("Error method is Empty :  %s", method)
	}
}
