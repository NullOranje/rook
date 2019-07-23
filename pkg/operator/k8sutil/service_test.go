package k8sutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
)

func TestGetServiceType(t *testing.T) {
	var result v1.ServiceType
	var err error

	result, err = GetServiceType("ClusterIP")
	assert.Equal(t, v1.ServiceTypeClusterIP, result)
	assert.Nil(t, err)

	result, err = GetServiceType("NodePort")
	assert.Equal(t, v1.ServiceTypeNodePort, result)
	assert.Nil(t, err)

	result, err = GetServiceType("LoadBalancer")
	assert.Equal(t, v1.ServiceTypeLoadBalancer, result)
	assert.Nil(t, err)

	_, err = GetServiceType("badString")
	assert.NotNil(t, err)
}
