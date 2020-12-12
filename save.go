package htmltojson

import (
	"encoding/json"
	"os"
)

// Save save a node to ./html_node.json
func Save(node *Node) error {
	return save(node, "./html_node.json")
}

// SaveNodes saves array of nodes to ./html_nodes.json
func SaveNodes(nodes []Node) error {
	return save(nodes, "./html_nodes.json")
}

// SaveToPath save a node to path
func SaveToPath(node *Node, path string) error {
	return save(node, path)
}

// SaveNodesToPath saves array of nodes to path
func SaveNodesToPath(nodes []Node, path string) error {
	return save(nodes, path)
}

func save(node interface{}, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewEncoder(f).Encode(node)
}
