package ipam

import (
	"fmt"
	"log"
	// "log"
	"net"
	"testing"

	"github.com/socketplane/socketplane/Godeps/_workspace/src/github.com/socketplane/ecc"
)

func TestInit(t *testing.T) {
	err := Init("eth1", true)
	if err != nil {
		t.Error("Error starting Consul ", err)
	}
}

func TestGetIpFullMask(t *testing.T) {
	count := 2500
	for i := 1; i < count; i++ {
		_, ipNet, _ := net.ParseCIDR("192.168.0.0/16")
		address := Request(*ipNet)
		address = address.To4()
		fmt.Println(address.String())
		if i%256 != int(address[3]) || i/256 != int(address[2]) {
			t.Error(address.String())
		}
	}
}

func TestGetIpPartialMask(t *testing.T) {
	count := 1000
	for i := 1; i < count; i++ {
		_, ipNet, _ := net.ParseCIDR("192.169.32.0/20")
		address := Request(*ipNet)
		address = address.To4()
		fmt.Println(address.String())
		if i%256 != int(address[3]) || 32+i/256 != int(address[2]) {
			t.Error(address.String())
		}
	}
}

func TestGetIpDeprecated(t *testing.T) {
	for i := 1; i < 250; i++ {
		addressStr, err := GetAnAddress("192.167.1.0/24")
		if err != nil {
			log.Println(err)
			t.Fatal(err)
		}
		address := net.ParseIP(addressStr).To4()
		if i != int(address[3]) {
			t.Error(addressStr)
		}
	}
}

func TestCleanup(t *testing.T) {
	ecc.Delete(dataStore, "192.167.1.0/24")
	ecc.Delete(dataStore, "192.168.0.0/16")
	ecc.Delete(dataStore, "192.169.32.0/20")
	Leave()
}
