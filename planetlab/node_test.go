package planetlab

import (
	"testing"

	"gotest.tools/assert"
)

// note that this test will fail if the node cse-yellow.cse.chalmers.se is removed
const existingNodeID = 2667
const existingHostname = "cse-yellow.cse.chalmers.se"
const nonExistingNodeID = -1
const nonExistingHostname = "foo.bar.baz"

func TestGetNodes(t *testing.T) {
	auth := GetTestAuth()
	nodeFilter := struct {
		NodeID int `xmlrpc:"node_id"`
	}{existingNodeID}

	// should return 1 node with matching ID
	nodes, err := GetNodes(auth, nodeFilter)
	assert.NilError(t, err)
	assert.Assert(t, len(nodes) == 1)
	assert.Assert(t, nodes[0].NodeID == existingNodeID)

	// should return 0 nodes since node with id -1 does not exist
	nodeFilter.NodeID = nonExistingNodeID
	nodes, err = GetNodes(auth, nodeFilter)
	assert.NilError(t, err)
	assert.Assert(t, len(nodes) == 0)
}

func TestGetNodeByID(t *testing.T) {
	auth := GetTestAuth()

	// should return node with matching ID
	node, err := GetNodeByID(auth, existingNodeID)
	assert.NilError(t, err)
	assert.Assert(t, node.NodeID == existingNodeID)

	// should return err, since node with id -1 does not exist
	node, err = GetNodeByID(auth, nonExistingNodeID)
	assert.Assert(t, err != nil)
	assert.Assert(t, node == nil)
}

func TestGetNodeByHostname(t *testing.T) {
	auth := GetTestAuth()

	// should return node with matching hostname
	node, err := GetNodeByHostName(auth, existingHostname)
	assert.NilError(t, err)
	assert.Assert(t, node.HostName == existingHostname)

	// should return err, since node with hostname foo.bar.baz does not exist
	node, err = GetNodeByHostName(auth, nonExistingHostname)
	assert.Assert(t, err != nil)
	assert.Assert(t, node == nil)
}
