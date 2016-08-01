package config

import (
	"testing"
)

func TestConfigFileMustBeLoaded(t *testing.T) {
	configuration := New()
	_, err := configuration.Load()
	if err != nil {
		t.Errorf("Error Config don't working :  %v", err)
	}

}

func TestIPMustBeNonEmpty(t *testing.T) {
	configuration := New()
	conf, _ := configuration.Load()
	ip := conf.Services[0].Ip
	if ip == "" {
		t.Errorf("Error ip is Empty :  %s", ip)
	}
}

func TestHostMustBeNonEmpty(t *testing.T) {
	configuration := New()
	conf, _ := configuration.Load()
	host := conf.Services[0].Host
	if host == "" {
		t.Errorf("Error host is Empty :  %s", host)
	}
}

func TestMethodMustBeNonEmpty(t *testing.T) {
	configuration := New()
	conf, _ := configuration.Load()
	method := conf.Services[0].Method
	if method == "" {
		t.Errorf("Error method is Empty :  %s", method)
	}
}
