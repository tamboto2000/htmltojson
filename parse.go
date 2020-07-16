package htmltojson

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Node struct {
	Type      string `json:"type"`
	Data      string `json:"data"`
	Namespace string `json:"namespace"`
	Attr      []Attr `json:"attr"`
	Child     []Node `json:"child"`
}

type Attr struct {
	Namespace string `json:"namespace"`
	Key       string `json:"key"`
	Val       string `json:"val"`
}

//node types
const (
	Text     = "text"
	Document = "document"
	Element  = "element"
	Comment  = "comment"
	Doctype  = "doctype"
)

//Parse parse HTML node to marshalable node
func Parse(root *html.Node) *Node {
	return parseToJSON(root)
}

//ParseString parse HTML string to marshalable node
func ParseString(str string) (*Node, error) {
	doc, err := html.Parse(strings.NewReader(str))
	if err != nil {
		return nil, err
	}

	return parseToJSON(doc), nil
}

//ParseBytes parse HTML bytes to marshalable node
func ParseBytes(byts []byte) (*Node, error) {
	doc, err := html.Parse(bytes.NewBuffer(byts))
	if err != nil {
		return nil, err
	}

	return parseToJSON(doc), nil
}

//ParseFromReader parse reader to marshalable node
func ParseFromReader(reader io.Reader) (*Node, error) {
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}

	return parseToJSON(doc), nil
}

//SearchNode search a node matched with options
func SearchNode(ty, data, namespace, key, val string, node *Node) *Node {
	return searchNode(ty, data, namespace, key, val, node)
}

//SearchAllNode search nodes matched with options
func SearchAllNode(ty, data, namespace, key, val string, node *Node) []Node {
	return searchAllNode(ty, data, namespace, key, val, node)
}

func searchAllNode(ty, data, namespace, key, val string, node *Node) []Node {
	nodes := make([]Node, 0)
	if isNodeMatch(ty, data, namespace, key, val, node) {
		nodes = append(nodes, *node)
	}

	if node.Child != nil {
		for _, child := range node.Child {

			nodes = append(nodes, searchAllNode(ty, data, namespace, key, val, &child)...)
		}
	}

	return nodes
}

func searchNode(ty, data, namespace, key, val string, node *Node) *Node {
	if isNodeMatch(ty, data, namespace, key, val, node) {
		return node
	}

	if node.Child != nil {
		for _, child := range node.Child {
			newNode := searchNode(ty, data, namespace, key, val, &child)
			if newNode != nil {
				return newNode
			}
		}
	}

	return nil
}

func isNodeMatch(ty, data, namespace, key, val string, node *Node) bool {
	if ty != "" && node.Type != ty {
		return false
	}

	if data != "" && node.Data != data {
		return false
	}

	if namespace != "" && node.Namespace != namespace {
		return false
	}

	if key != "" {
		found := false
		for _, attr := range node.Attr {
			if attr.Key == key {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	if val != "" {
		found := false
		for _, attr := range node.Attr {
			if attr.Val == val {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

func parseToJSON(root *html.Node) *Node {
	newNode := new(Node)

	if root.Type == html.TextNode {
		newNode.Type = Text
	} else if root.Type == html.DocumentNode {
		newNode.Type = Document
	} else if root.Type == html.ElementNode {
		newNode.Type = Element
	} else if root.Type == html.CommentNode {
		newNode.Type = Comment
	} else if root.Type == html.DoctypeNode {
		newNode.Type = Doctype
	}

	newNode.Data = root.Data
	newNode.Namespace = root.Namespace

	for _, attr := range root.Attr {
		newNode.Attr = append(newNode.Attr, Attr{
			Namespace: attr.Namespace,
			Key:       attr.Key,
			Val:       attr.Val,
		})
	}

	for c := root.FirstChild; c != nil; c = c.NextSibling {
		if result := parseToJSON(c); result != nil {
			newNode.Child = append(newNode.Child, *result)
		}
	}

	return newNode
}
