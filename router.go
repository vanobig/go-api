package main

import (
	"net/http"
	"net/url"
	"strings"
)

type Router struct {
	tree        *node
	rootHandler Handle
}

type Handle func(http.ResponseWriter, *http.Request, url.Values)

// Represents a struct of each node in the tree.
type node struct {
	children     []*node
	component    string
	isNamedParam bool
	methods      map[string]Handle
}

func NewRouter(rootHandler Handle) *Router {
	node := node{
		component: "/",
		isNamedParam: false,
		methods: make(map[string]Handle)}

	// Automatically add listed routes
	for _, route := range routes {
		// Makes sure our path is valid
		if route.Pattern[0] != '/' {
			panic("Path has to start with a /")
		}

		var handler Handle

		handler = route.HandlerFunc
		handler = Logger(handler)

		node.addNode(route.Method, route.Pattern, handler)
	}

	return &Router{tree: &node, rootHandler: rootHandler}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	params := req.Form
	node, _ := r.tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)

	if handler := node.methods[req.Method]; handler != nil {
		handler(w, req, params)
	} else {
		r.rootHandler(w, req, params)
	}
}

func (n *node) addNode(method, path string, handler Handle) {
	components := strings.Split(path, "/")[1:]
	count := len(components)

	for {
		aNode, component := n.traverse(components, nil)

		// update an existing node
		if aNode.component == component && count == 1 {
			aNode.methods[method] = handler
			return
		}

		newNode := node{
			component: component,
			isNamedParam: false,
			methods: make(map[string]Handle)}

		// check if it is a named param
		if len(component) > 0 && component[0] == ':' {
			newNode.isNamedParam = true
		}

		// this is the last component of the url resource, so it gets the handler
		if count == 1 {
			newNode.methods[method] = handler
		}

		aNode.children = append(aNode.children, &newNode)
		count--

		if count == 0 {
			break
		}
	}
}

func (n *node) traverse(components []string, params url.Values) (*node, string) {
	component := components[0]

	if len(n.children) > 0 {
		for _, child := range n.children {
			if component == child.component || child.isNamedParam {
				if child.isNamedParam && params != nil {
					params.Add(child.component[1:], component)
				}

				next := components[1:]

				if len(next) > 0 {
					return child.traverse(next, params) // recursion
				} else {
					return child, component
				}
			}
		}
	}

	return n, component
}