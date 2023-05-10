# ТГ бот

В этом проекте реализован обычный телеграм бот с ипользованием Telegram Bot Api

***Описание бота***

1) Приветственное сообщение
![Приветственное сообщение](https://github.com/verbovyar/VkBot/blob/main/images/hello.PNG)


2) Сам функционал бота

    ![Функционал](https://github.com/verbovyar/VkBot/blob/main/images/functional.PNG)

Вся работа происходит над некой сущностью - это наш объект хранения

Сама сущность
``` Golang
type Entity struct {
	Name string
	Age  uint
	Id   uint
}
``` 
Имеет три поля: имя, позраст и ключ.

*Теперь по порядку о каждой функции*
1) Функция add добавляет нашу сущность в базу данных.
2) Функция update позволяет поменять данные в нашей сущности, если произошли какие-то изменения.
3) Функция delete просто удаляет объект из базы данных.
4) Функция list позволяет вывести все данные, которые лежат в нашей БД.
5) Функция help позволяет вывести информацию о том, какие аргументы нужны для наших команд.

Все функции можно использовать при нажатии кнопочки.

В проекте я попытался воспользоваться приципом чистой архитектуры, чтобы в дальнейшем можно было
развивать проект и проще было изменять или внедрять какие-то части проекта. Весь проект разбит
на смысловые части (слои).

![архитектура](https://github.com/verbovyar/VkBot/blob/main/images/architecture.PNG).

Для базы данных спользован интрефейс, чтобы в дальнейшем можно было беспрепядственно подключать
разные виды БД и не иметь зависимости от других частей проекта. Можно спокойно имплементировать
интерфейс и не менять хендлеры для другой БД.

Так же отдельно вынесен в отдельную часть конфиг файл и для него написана своя функция.
```Golang
type Config struct {
ApiKey string `mapstructure:"API_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
viper.AddConfigPath(path)
viper.SetConfigName("conf")
viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		_ = fmt.Errorf("do not parse config file:%v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		_ = fmt.Errorf("do not parse config file:%v", err)
	}

	return
}
```
