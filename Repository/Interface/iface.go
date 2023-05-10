package Interface

import (
	"bot/Domain"
)

type Repository interface {
	Add(player *Domain.Entity) error
	List() []*Domain.Entity
	Update(user *Domain.Entity, id uint) error
	Delete(id uint) error
}
