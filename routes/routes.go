package routes

import (
	"crud-movies/config"
	"crud-movies/controllers"
	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {
	//init echo
	e := echo.New()

	//Database
	db := config.CreateConnection()

	//movie routes
	e.GET("/movies", controllers.GetMovies(db))
	e.POST("/movies", controllers.AddMovie(db))
	e.PUT("/movies/:id", controllers.UpdateMovie(db))
	e.DELETE("/movies/:id", controllers.DeleteMovie(db))

	return e
}