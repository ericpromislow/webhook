package integration_test

import (
	"encoding/json"
	"github.com/stretchr/testify/require"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
)

func (m *IntegrationSuite) TestManagementCluster() {
	newObj := func() *v3.Cluster { return &v3.Cluster{} }

	var validCreateObj v3.Cluster
	// Use a json-encoded representation of the desired data
	// to avoid loading the RKE types module.
	s := `{
  "metadata": {
    "name": "test-cluster"
  },
  "spec": {
    "rancherKubernetesEngineConfig": { }
  }
}`
	err := json.Unmarshal([]byte(s), &validCreateObj)
	require.NoError(m.T(), err, "Failed to create v3.Cluster object from json")

	validDelete := func() *v3.Cluster {
		return &validCreateObj
	}
	endPoints := &endPointObjs[*v3.Cluster]{
		invalidCreate:  nil,
		newObj:         newObj,
		validCreateObj: &validCreateObj,
		invalidUpdate:  nil,
		validUpdate:    nil,
		validDelete:    validDelete,
	}
	validateEndpoints(m.T(), endPoints, m.clientFactory)
}
