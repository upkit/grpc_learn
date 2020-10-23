package dao

import (
	"context"
	"time"

	"github.com/upkit/grpc_learn/model"
	"github.com/upkit/grpc_learn/proto/person/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (d *Dao) checkAPI(api string) error {
	if api != model.APIVersion {
		msg := "unsupported API version: servive implements API version " + model.APIVersion
		return status.Error(codes.Unimplemented, msg)
	}
	return nil
}

func (d *Dao) Create(ctx context.Context, request *person.CreateRequest) (*person.CreateResponse, error) {
	err := d.checkAPI(request.Api)
	if err != nil {
		return nil, err
	}

	request.Person.Id = uint64(time.Now().UnixNano() / 1e3)
	_, err = d.dc.NamedExecContext(ctx, "INSERT INTO tb_person (id, name, age) VALUE (:id, :name, :age)", request.Person)
	if err != nil {
		return nil, err
	}

	return &person.CreateResponse{Api: model.APIVersion, UserId: request.Person.Id}, nil
}

func (d *Dao) Read(ctx context.Context, request *person.ReadRequest) (*person.ReadResponse, error) {
	err := d.checkAPI(request.Api)
	if err != nil {
		return nil, err
	}

	var p person.Person
	err = d.dc.GetContext(ctx, &p, "SELECT id, name, age FROM tb_person WHERE id=?", request.UserId)
	if err != nil {
		return nil, err
	}

	return &person.ReadResponse{Api: model.APIVersion, Person: &p}, nil
}

func (d *Dao) Update(ctx context.Context, request *person.UpdateRequest) (*person.UpdateResponse, error) {
	err := d.checkAPI(request.Api)
	if err != nil {
		return nil, err
	}

	_, err = d.dc.ExecContext(ctx, "UPDATE tb_person SET name=?, age=? WHERE id=?", request.Person.Name, request.Person.Age, request.Person.Id)
	if err != nil {
		return nil, err
	}

	return &person.UpdateResponse{Api: model.APIVersion, ChangedId: request.Person.Id}, nil
}

func (d *Dao) Delete(ctx context.Context, request *person.DeleteRequest) (*person.DeleteResponse, error) {
	err := d.checkAPI(request.Api)
	if err != nil {
		return nil, err
	}

	_, err = d.dc.ExecContext(ctx, "DELETE FROM tb_person WHERE id=?", request.UserId)
	if err != nil {
		return nil, err
	}

	return &person.DeleteResponse{Api: model.APIVersion, DeletedId: request.UserId}, nil
}

func (d *Dao) ReadAll(ctx context.Context, request *person.ReadAllRequest) (*person.ReadAllResponse, error) {
	err := d.checkAPI(request.Api)
	if err != nil {
		return nil, err
	}

	var list []*person.Person
	err = d.dc.SelectContext(ctx, &list, "SELECT id, name, age FROM tb_person")
	if err != nil {
		return nil, err
	}

	return &person.ReadAllResponse{Api: model.APIVersion, Persons: list}, nil
}
