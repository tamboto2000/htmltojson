// Package htmltojson is a HTML parser, based on net/html package. This package actually just to simplify HTML parsing.
// If you need more complex HTML processing, please use net/html as its offer more features.
// The package name is actually is not really fitting for this package purpose, but I use this package for may scraper engines, so
// I don't really want to bother with changing the package name...
package htmltojson

import (
	"bytes"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Node is parsed HTML object
type Node struct {
	Type      string `json:"type"`
	Data      string `json:"data"`
	Namespace string `json:"namespace"`
	Attr      []Attr `json:"attr"`
	Child     []Node `json:"child"`
}

// Attr is HTML attributes, like class, style, id, etc.
type Attr struct {
	Namespace string `json:"namespace"`
	Key       string `json:"key"`
	Val       string `json:"val"`
}

// Node Types
const (
	Text     = "text"
	Document = "document"
	Element  = "element"
	Comment  = "comment"
	Doctype  = "doctype"
)

// Parse parse HTML node to marshalable node
func Parse(root *html.Node) *Node {
	return parseToJSON(root)
}

// ParseString parse HTML string to marshalable node
func ParseString(str string) (*Node, error) {
	doc, err := html.Parse(strings.NewReader(str))
	if err != nil {
		return nil, err
	}

	return parseToJSON(doc), nil
}

// ParseBytes parse HTML bytes to marshalable node
func ParseBytes(byts []byte) (*Node, error) {
	doc, err := html.Parse(bytes.NewBuffer(byts))
	if err != nil {
		return nil, err
	}

	return parseToJSON(doc), nil
}

// ParseFromReader parse reader to marshalable node
func ParseFromReader(reader io.Reader) (*Node, error) {
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}

	return parseToJSON(doc), nil
}

// ParseFromFile parse HTML from file in path
func ParseFromFile(path string) (*Node, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	return ParseFromReader(f)
}

// SearchNode search a node matched with params.
// ty for HTML object type,
// data is for HTML tag name,
// key is for attribute key
// val is for attribute value with key
func SearchNode(ty, data, namespace, key, val string, node *Node) *Node {
	return searchNode(ty, data, namespace, key, val, node)
}

// SearchAllNode search nodes matched with options.
// ty for HTML object type,
// data is for HTML tag name,
// key is for attribute key
// val is for attribute value with key
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
