package DataBase

import (
	"bot/Domain"
	"bot/Repository/Interface"
	"strconv"

	"github.com/pkg/errors"
)

var Map *Repository

func New() *Repository {
	Map = &Repository{
		mapData: make(map[uint]*Domain.Entity),
	}
	return Map
}

//--------Methods

type Repository struct { // IRepository
	mapData map[uint]*Domain.Entity

	Iface Interface.Repository
}

func (r *Repository) List() []*Domain.Entity {
	res := make([]*Domain.Entity, 0)

	for _, value := range Map.mapData {
		temp := &Domain.Entity{
			Name: value.Name,
			Age:  value.Age,
			Id:   value.Id}

		res = append(res, temp)
	}

	return res
}

func (r *Repository) Add(entity *Domain.Entity) error {
	if _, ok := Map.mapData[entity.GetId()]; ok {
		return errors.Wrap(errors.New("user exists"), strconv.FormatUint(uint64(entity.GetId()), 10))
	}
	Map.mapData[entity.GetId()] = entity

	return nil
}

func (r *Repository) Update(entity *Domain.Entity, id uint) error {
	if _, ok := Map.mapData[id]; !ok {
		return errors.Wrap(errors.New("user does not exist"), strconv.FormatUint(uint64(entity.GetId()), 10))
	}
	Map.mapData[id] = entity
	entity.Id = id

	return nil
}

func (r *Repository) Delete(id uint) error {
	if _, ok := Map.mapData[id]; ok {
		delete(Map.mapData, id)

		return nil
	}

	return errors.Wrap(errors.New("user does not exist"), strconv.FormatUint(uint64(id), 10))
}
