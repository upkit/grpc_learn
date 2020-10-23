package client

import (
	"testing"

	"github.com/upkit/grpc_learn/libs/conf"
	"github.com/upkit/grpc_learn/libs/log"
	"github.com/upkit/grpc_learn/model"
	"github.com/upkit/grpc_learn/proto/person/v1"
)

func defaultClient() *Client {
	log.Init(log.Conf{
		Path:  "client_test",
		Level: "debug",
	})
	return New(&conf.Config{
		Debug:    true,
		GrpcAddr: "127.0.0.1:8080",
	})
}

func TestClient_Create(t *testing.T) {
	c := defaultClient()
	defer c.Close()

	resp, err := c.Create(&person.CreateRequest{Api: model.APIVersion, Person: &person.Person{
		Name: "哆啦a梦",
		Age:  60,
	}})
	if err != nil {
		t.Errorf("create error, err:%v", err)
		return
	}
	t.Logf("create success, resp.id:%d", resp.UserId)
}

func TestClient_Read(t *testing.T) {
	c := defaultClient()
	defer c.Close()

	resp, err := c.Read(&person.ReadRequest{Api: model.APIVersion, UserId: 1603189466836641000})
	if err != nil {
		t.Errorf("read error, err:%v", err)
		return
	}
	t.Logf("read success, resp.p-> id:%d, name:%s, age:%d", resp.Person.Id, resp.Person.Name, resp.Person.Age)
}

func TestClient_Update(t *testing.T) {
	c := defaultClient()
	defer c.Close()

	resp, err := c.Update(&person.UpdateRequest{Api: model.APIVersion, Person: &person.Person{
		Id:   1603400369649694,
		Name: "静香",
		Age:  25,
	}})
	if err != nil {
		t.Errorf("update error, err:%v", err)
		return
	}
	t.Logf("update success, resp.id:%d", resp.ChangedId)
}

func TestClient_ReadAll(t *testing.T) {
	c := defaultClient()
	defer c.Close()

	resp, err := c.ReadAll(&person.ReadAllRequest{Api: model.APIVersion})
	if err != nil {
		t.Errorf("readAll error, err:%v", err)
		return
	}
	for i := range resp.Persons {
		t.Logf("readAll success, resp.list-> id:%d, name:%s, age:%d", resp.Persons[i].Id, resp.Persons[i].Name, resp.Persons[i].Age)
	}
}

func TestClient_Delete(t *testing.T) {
	c := defaultClient()
	defer c.Close()

	resp, err := c.Delete(&person.DeleteRequest{Api: model.APIVersion, UserId: 1603189466836641})
	if err != nil {
		t.Errorf("delete error, err:%v", err)
		return
	}
	t.Logf("delete success, resp.id:%d", resp.DeletedId)
}
