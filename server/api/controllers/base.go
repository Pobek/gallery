package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pobek/gallery/server/api/middlewares"
	"github.com/pobek/gallery/server/api/models"
	"github.com/pobek/gallery/server/api/responses"
)

// App - the structure of the application object
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Init - initializes the application
func (app *App) Init(DbHost, DbPort, DbUser, DbName, DbPassword string) {
	var err error
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		DbHost,
		DbPort,
		DbUser,
		DbName,
		DbPassword)

	app.DB, err = gorm.Open("postgres", DBURI)
	if err != nil {
		log.Fatalf("Cannot connect to db '%s' on host '%s'. Database connection error: %v", DbName, DbHost, err)
	} else {
		log.Printf("Connected to database '%s' on host '%s'", DbName, DbHost)
	}

	app.DB.Debug().AutoMigrate(&models.User{})

	app.Router = mux.NewRouter().StrictSlash(true)
	app.initRoutes()
}

func (app *App) initRoutes() {
	app.Router.Use(middlewares.SetContentTypeMiddleware) // sets content-type to json

	app.Router.HandleFunc("/", home).Methods("GET")
	app.Router.HandleFunc("/register", app.UserSignUp).Methods("POST")
	app.Router.HandleFunc("/login", app.Login).Methods("POST")
}

// RunServer - starts the server
func (app *App) RunServer() {
	log.Printf("Server starting on port 5000")
	log.Fatal(http.ListenAndServe(":5000", app.Router))
}

func home(respWriter http.ResponseWriter, req *http.Request) {
	responses.JSON(respWriter, http.StatusOK, "Welcome To GalleryAPI")
}
