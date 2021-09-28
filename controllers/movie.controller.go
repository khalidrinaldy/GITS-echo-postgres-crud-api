package controllers

import (
	"crud-movies/models"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetMovies(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		sqlQuery := "SELECT id, title, rating from movies order by id"
		rows, err := db.Query(sqlQuery)
		if err!=nil {
			fmt.Println(err)
		}
		defer rows.Close()
		result := models.Movies{}

		for rows.Next() {
			movie := models.Movie{}
			err2 := rows.Scan(&movie.Id, &movie.Title, &movie.Rating)
			if err2!=nil {
				return err2
			}
			result.Movies = append(result.Movies, movie)
		}
		return c.JSON(http.StatusCreated, result)
	}
}

func AddMovie(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		movie := models.Movie{}
		if err:= c.Bind(&movie); err!=nil {
			return err
		}
		movie.Title = c.FormValue("title")
		movie.Rating, _ = strconv.Atoi(c.FormValue("rating"))

		sqlQuery := "INSERT INTO movies (title, rating) VALUES ($1, $2)"
		res, err := db.Query(sqlQuery, movie.Title, movie.Rating)
		if err!=nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, movie)
		}
		return c.String(http.StatusOK, "ok")
	}
}

func UpdateMovie(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		movie := models.Movie{}
		if err:= c.Bind(&movie); err!=nil {
			return err
		}
		id := c.Param("id")
		fmt.Println(id)
		movie.Title = c.FormValue("title")
		movie.Rating, _ = strconv.Atoi(c.FormValue("rating"))

		sqlQuery := "UPDATE movies SET title=$1, rating=$2 WHERE id=$3"
		res, err := db.Query(sqlQuery, movie.Title, movie.Rating, id)
		if err!=nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, movie)
		}
		return c.String(http.StatusOK, id + " Updated")
	}
}

func DeleteMovie(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		sqlQuery := "DELETE FROM movies WHERE id=$1"
		res,err := db.Query(sqlQuery, id)
		if err!=nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, "Deleted")
		}
		return c.String(http.StatusOK, id + " Deleted")
	}
}