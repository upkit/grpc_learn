package dao

import (
	"context"
	"testing"
	"time"

	"github.com/upkit/grpc_learn/libs/cache/mredis"
	"github.com/upkit/grpc_learn/libs/conf"
	"github.com/upkit/grpc_learn/libs/db/msql"
	"github.com/upkit/grpc_learn/libs/log"
	"github.com/upkit/grpc_learn/model"
	"github.com/upkit/grpc_learn/proto/person/v1"
)

func defaultClient() *Dao {
	log.Init(log.Conf{
		Path:  "server_test",
		Level: "debug",
	})
	return New(context.Background(), &conf.Config{
		Debug: true,
		Mysql: msql.Conf{
			Addr: "root:1234@tcp(127.0.0.1:3307)/test?charset=utf8&parseTime=true&loc=Local",
			Open: 20,
			Idle: 10,
			Life: time.Hour,
		},
		Redis: mredis.Conf{
			Addr:       "127.0.0.1:6379",
			DB:         0,
			Pool:       10,
			MaxRetries: 3,
		},
	})
}

func TestPersonServer_Create(t *testing.T) {
	s := defaultClient()

	resp, err := s.Create(context.Background(), &person.CreateRequest{Api: model.APIVersion, Person: &person.Person{
		Name: "小哞",
		Age:  25,
	}})
	if err != nil {
		t.Errorf("create failed, err:%v", err)
		return
	}
	t.Logf("user_id:%d", resp.UserId)
}

func TestPersonServer_Read(t *testing.T) {
	s := defaultClient()

	resp, err := s.Read(context.Background(), &person.ReadRequest{Api: model.APIVersion, UserId: 1603189431803871000})
	if err != nil {
		t.Errorf("read failed, err:%v", err)
		return
	}
	t.Logf("person:%+v", resp.Person)
}

func TestPersonServer_Update(t *testing.T) {
	s := defaultClient()

	resp, err := s.Update(context.Background(), &person.UpdateRequest{Api: model.APIVersion, Person: &person.Person{
		Id:   1603189431803871000,
		Name: "大猩猩",
		Age:  34,
	}})
	if err != nil {
		t.Errorf("update failed, err:%v", err)
		return
	}
	t.Logf("changed_id:%d", resp.ChangedId)
}

func TestPersonServer_Delete(t *testing.T) {
	s := defaultClient()

	resp, err := s.Delete(context.Background(), &person.DeleteRequest{Api: model.APIVersion, UserId: 1603189431803871000})
	if err != nil {
		t.Errorf("delete failed, err:%v", err)
		return
	}
	t.Logf("deleted_id:%d", resp.DeletedId)
}

func TestPersonServer_ReadAll(t *testing.T) {
	s := defaultClient()

	resp, err := s.ReadAll(context.Background(), &person.ReadAllRequest{Api: model.APIVersion})
	if err != nil {
		t.Errorf("readAll failed, err:%v", err)
		return
	}

	t.Logf("persons:%+v", resp.Persons)
}
