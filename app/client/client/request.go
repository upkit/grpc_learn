package client

import (
	"github.com/upkit/grpc_learn/proto/person/v1"
)

func (c *Client) Create(request *person.CreateRequest) (*person.CreateResponse, error) {
	return c.cc.Create(c.ctx, request)
}

func (c *Client) Read(request *person.ReadRequest) (*person.ReadResponse, error) {
	return c.cc.Read(c.ctx, request)
}

func (c *Client) Update(request *person.UpdateRequest) (*person.UpdateResponse, error) {
	return c.cc.Update(c.ctx, request)
}

func (c *Client) Delete(request *person.DeleteRequest) (*person.DeleteResponse, error) {
	return c.cc.Delete(c.ctx, request)
}

func (c *Client) ReadAll(request *person.ReadAllRequest) (*person.ReadAllResponse, error) {
	return c.cc.ReadAll(c.ctx, request)
}
