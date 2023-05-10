package Service

import (
	"bot/Domain"
	"bot/Repository/DataBase"
	"bot/Repository/Interface"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

const (
	helpCmd   = "help"
	listCmd   = "list"
	addCmd    = "add"
	changeCmd = "change"
	deleteCmd = "delete"
)

var Service *UserService

func New(repository Interface.Repository) *UserService {
	Service = &UserService{
		repository: repository,
	}
	return Service
}

type UserService struct {
	repository Interface.Repository
}

//---------

func listFunc(string) string {
	data := DataBase.Map.List()
	res := make([]string, 0, len(data))

	for _, value := range data {
		res = append(res, value.GetData())
	}

	return strings.Join(res, "\n")
}

func helpFunc(string) string {
	return "help - list commands\n" +
		"list - list data\n" +
		"add <name> <age> - add new person \n" +
		"change <id> <name> <age> - change info about person\n" +
		"delete <id> - delete person"
}

func addFunc(data string) string {
	params := strings.Split(data, " ")
	if len(params) != 2 {
		return errors.Wrapf(errors.New("bad argument"), "%d items: <%v>", len(params), params).Error()
	}

	age, _ := strconv.Atoi(params[1])
	user, err := Domain.NewEntity(params[0], uint(age))
	if err != nil {
		return err.Error()
	}

	err = DataBase.Map.Add(user)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("user (%d) %s:%d added", user.Id, user.Name, user.Age)
}

func changeFunc(data string) string {
	params := strings.Split(data, " ")
	if len(params) != 3 {
		return errors.Wrapf(errors.New("bad argument"), "%d items: <%v>", len(params), params).Error()
	}

	age, _ := strconv.Atoi(params[2])

	user := Domain.Entity{}

	if err := user.SetName(params[1]); err != nil {
		fmt.Println(err)
		return ""
	}

	if err := user.SetAge(uint(age)); err != nil {
		fmt.Println(err)
		return ""
	}

	id, _ := strconv.Atoi(params[0])
	err := DataBase.Map.Update(&user, uint(id))
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("user info was changed")
}

func deleteFunc(data string) string {
	params := strings.Split(data, "\n")
	if len(params) != 1 {
		return errors.Wrapf(errors.New("bad argument"), "%d items: <%v>", len(params), params).Error()
	}

	id, _ := strconv.Atoi(params[0])

	err := DataBase.Map.Delete(uint(id))
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("user deleted")
}

func AddHandlers(bot *Bot) {
	bot.RegisterHandler(helpCmd, helpFunc)
	bot.RegisterHandler(listCmd, listFunc)
	bot.RegisterHandler(addCmd, addFunc)
	bot.RegisterHandler(changeCmd, changeFunc)
	bot.RegisterHandler(deleteCmd, deleteFunc)
}
