package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetK8sNamespaces(t *testing.T) {
	actual := GetK8sNamespaces()

	assert.Contains(t, actual, "kube-system")
}

func TestGetConGroups(t *testing.T) {
	actual := GetConGroups("kube-system")

	for _, group := range actual {
		if group.MicroserviceName != "kube-system" {
			t.Errorf("it should have %s namespace", "kube-system")
		}
	}
}

func TestGetServices(t *testing.T) {
	actual := GetServices("kube-system")

	for _, svc := range actual {
		if svc.MicroserviceName != "kube-system" {
			t.Errorf("it should have %s namespace", "kube-system")
		}
	}
}

func TestGetEndpoints(t *testing.T) {
	actual := GetEndpoints("kube-system")

	for _, endpoint := range actual {
		if endpoint.MicroserviceName != "kube-system" {
			t.Errorf("it should have %s namespace", "kube-system")
		}
	}
}
