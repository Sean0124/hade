package framework

import "net/http"

// framework core struct
type Core struct {
}

// Initialize a framework core struct
func NewCore() *Core {
	return &Core{}
}

// framework core struct implementation Handler interface
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {

}
