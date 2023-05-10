package Domain

import "fmt"

type Entity struct {
	Name string
	Age  uint
	Id   uint
}

var LastId = uint(0)

func NewEntity(name string, age uint) (*Entity, error) {
	user := Entity{}

	if err := user.SetName(name); err != nil {
		return nil, err
	}

	if err := user.SetAge(age); err != nil {
		return nil, err
	}

	user.Id = LastId
	LastId++

	return &user, nil
}

func (user *Entity) SetName(name string) error {
	nameSize := len(name)
	if nameSize == 0 || nameSize > 10 {
		return fmt.Errorf("bad name <%v>", name)
	}

	user.Name = name

	return nil
}

func (user *Entity) SetAge(age uint) error {
	if age < 0 {
		return fmt.Errorf("bad age <%v>", age)
	}

	user.Age = age

	return nil
}

func (user *Entity) GetName() string {
	return user.Name
}

func (user *Entity) GetAge() uint {
	return user.Age
}

func (user *Entity) GetData() string {
	return fmt.Sprintf("(%d) %s: %d", user.Id, user.Name, user.Age)
}

func (user *Entity) GetId() uint {
	return user.Id
}
