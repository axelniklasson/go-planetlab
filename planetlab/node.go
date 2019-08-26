package planetlab

import "fmt"

// Node represents a PlanetLab node
type Node struct {
	LastUpdated       int    `xmlrpc:"last_updated"`
	BootState         string `xmlrpc:"boot_state"`
	SiteID            int    `xmlrpc:"site_id"`
	PcuIDs            []int  `xmlrpc:"pcu_ids"`
	Session           string `xmlrpc:"session"`
	SSHRSAKey         string `xmlrpc:"ssh_rsa_key"`
	LastContact       int    `xmlrpc:"last_contact"`
	PeerNodeID        int    `xmlrpc:"peer_node_id"`
	HostName          string `xmlrpc:"hostname"`
	SliceIDs          []int  `xmlrpc:"slice_ids"`
	Version           string `xmlrpc:"version"`
	PeerID            int    `xmlrpc:"peer_id"`
	NodeID            int    `xmlrpc:"node_id"`
	Key               string `xmlrpc:"key"`
	ConfFileIDs       []int  `xmlrpc:"conf_file_ids"`
	NodegroupIDs      []int  `xmlrpc:"nodegroup_ids"`
	SliceIDsWhitelist []int  `xmlrpc:"slice_ids_whitelist"`
	NodenetworkIDs    []int  `xmlrpc:"nodenetwork_ids"`
	BootNonce         string `xmlrpc:"boot_nonce"`
	DateCreated       int    `xmlrpc:"date_created"`
	Model             string `xmlrpc:"model"`
	Ports             []int  `xmlrpc:"ports"`
}

func (n Node) String() string {
	return fmt.Sprintf("Node - NodeID: %d, Hostname: %s Boot state: %s", n.NodeID, n.HostName, n.BootState)
}

// GetNodes returns a slice of Nodes containing details about nodes
//
// URL: https://www.planet-lab.org/doc/plc_api#GetNodes
func GetNodes(auth Auth, nodeFilter interface{}) ([]Node, error) {
	client := GetClient()
	args := make([]interface{}, 2)
	args[0] = auth
	args[1] = nodeFilter
	nodes := []Node{}

	err := client.Call(getNodesMethod, args, &nodes)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

// GetNodeByID is a wrapper around GetNodes that makes it easier to fetch a node using its ID
func GetNodeByID(auth Auth, nodeID int) (*Node, error) {
	nodeFilter := struct {
		NodeID int `xmlrpc:"node_id"`
	}{nodeID}

	nodes, err := GetNodes(auth, nodeFilter)
	if err != nil {
		return nil, err
	} else if len(nodes) != 1 {
		return nil, fmt.Errorf("Got %d nodes matching filter %v", len(nodes), nodeFilter)
	}

	return &nodes[0], nil
}

// GetNodeByHostName is a wrapper around GetNodes that makes it easier to fetch a node using its hostname
func GetNodeByHostName(auth Auth, hostname string) (*Node, error) {
	nodeFilter := struct {
		Hostname string `xmlrpc:"hostname"`
	}{hostname}

	nodes, err := GetNodes(auth, nodeFilter)
	if err != nil {
		return nil, err
	} else if len(nodes) != 1 {
		return nil, fmt.Errorf("Got %d nodes matching filter %v", len(nodes), nodeFilter)
	}

	return &nodes[0], nil
}
