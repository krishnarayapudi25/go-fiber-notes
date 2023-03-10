package router


import(

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    noteRoutes "github.com/krishnarayapudi25/go-fiber-notes/internals/routes/note"

)
    

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())

	noteRoutes.SetupNoteRoutes(api)

}