package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type Book struct {
	gorm.Model
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
}

func InitialMigration() {
	DB.AutoMigrate(&Book{})
}

// database connection
func InitDB() {

	// declare struct config & variable connectionString
	connectionString := "root:qwerty123@tcp(127.0.0.1:3306)/immersive?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func init() {
	InitDB()
	InitialMigration()
}

func main() {
	// create a new echo instance
	e := echo.New()

	// Route / to handler function
	e.GET("/books", GetBooksController)
	e.GET("/books/:id", GetOneBookController)
	e.POST("/books", CreateBookController)

	// start the server, and log if it fails
	e.Start(":8000")
}

// get all users
func GetOneBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	var book []Book

	if err := DB.Find(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"book":    book,
	})
}

// get all users
func GetBooksController(c echo.Context) error {
	var books []Book

	if err := DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// create new user
func CreateBookController(c echo.Context) error {
	book := Book{}
	c.Bind(&book)

	if err := DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// func main() {
// 	fmt.Println("Hello Main")
// }

// func init() {
// 	fmt.Println("Hello Init")
// }
