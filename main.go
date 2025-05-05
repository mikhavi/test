package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Подключаем HTML-шаблонизатор
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Статика (CSS и прочее)
	app.Static("/static", "./public", fiber.Static{
		Browse: true,
	})

	// Главная страница
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Главная",
		})
	})

	// Страница "О сайте"
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{
			"Title": "О сайте",
		})
	})

	app.Listen(":3000")
}
