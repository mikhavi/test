package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func main() {
	// Инициализируем шаблонизатор
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Подключение статики (CSS и т.д.)
	app.Static("/static", "./public")

	// Подключение к базе данных
	user := "admin_user"
	password := "Test123!"
	server := "localhost"
	port := "1433"
	database := "bookstore"

	connStr := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, password, server, port, database)

	var err error
	db, err = gorm.Open(sqlserver.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	// Главная страница (форма логина)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Форма входа",
		})
	})

	// Обработка логина (POST)
	app.Post("/login", func(c *fiber.Ctx) error {
		type Credentials struct {
			Username string `form:"username"`
			Password string `form:"password"`
		}

		var creds Credentials
		if err := c.BodyParser(&creds); err != nil {
			return c.SendString("Ошибка разбора формы")
		}

		// Пример проверки: можно заменить на запрос в базу данных
		if creds.Username == "admin" && creds.Password == "1234" {
			return c.SendString("Добро пожаловать, " + creds.Username)
		}

		return c.SendString("Неверный логин или пароль")
	})

	// О сайте
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{
			"Title": "О сайте",
		})
	})

	log.Println("Сервер запущен на http://localhost:3000")
	app.Listen(":3000")
}
