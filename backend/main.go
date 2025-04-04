package main

import (
	"log"
	"net/http"
	"note_server/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
	err error
	notes = []models.Note{
		{Title: "Test", Description: "Very interesting..."},
	}
)


func getNotes(c *gin.Context){
	notes := []models.Note{}
	db.Find(&notes)
	c.JSON(200, &notes)
}

func getNote(c *gin.Context){
	var note models.Note

	db.Find(&note, c.Param("id"))

	if db.Error != nil {
		c.String(http.StatusNotFound, "Error: %s" , c.Error(db.Error))
		return
	}

	c.JSON(http.StatusOK, &note)
}

func createNote(c *gin.Context){
	var note models.Note

	if err := c.BindJSON(&note); err!=nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if db.Create(&note); db.Error!=nil {
		c.String(http.StatusBadGateway, db.Error.Error())
		return
	}

	c.JSON(http.StatusOK, &note)
}

func updateNote(c *gin.Context){
	var note models.Note
	
	id_param, _ := strconv.Atoi(c.Param("id"))

	note.ID = uint(id_param)
	
	transaction := db.Begin()
	
	if transaction.Find(&note); transaction.Error!=nil {
		c.String(http.StatusNotFound, transaction.Error.Error())
		transaction.Rollback()
		return
	}

	if err := c.BindJSON(&note); err!=nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else if note.ID != uint(id_param) {
		c.String(http.StatusBadRequest, 
			"Id in the parameter and in the body don't match. Expected: %s, got: %s", id_param, note.ID)
		return
	}

	if transaction.Save(&note); transaction.Error!=nil {
		c.String(http.StatusBadGateway, transaction.Error.Error())
		transaction.Rollback()
		return
	}

	transaction.Commit()
	
	c.JSON(http.StatusOK, &note)
}

func deleteNote(c *gin.Context){
	var note models.Note

	id_param, _ := strconv.Atoi(c.Param("id"))
	
	note.ID = uint(id_param)

	transaction := db.Begin()
	
	if transaction.Find(&note); transaction.Error!=nil {
		c.String(http.StatusNotFound, transaction.Error.Error())
		transaction.Rollback()
		return
	}

	if transaction.Delete(&note); transaction.Error!=nil {
		c.String(http.StatusBadGateway, transaction.Error.Error())
		transaction.Rollback()
		return
	}

	if db := transaction.Commit(); db.Error != nil {
		c.String(http.StatusInternalServerError, db.Error.Error())
		return
	}

	c.JSON(http.StatusOK, &note)
}

func SetupServer() (*gin.Engine) {
	log.Print("Setting up server")
	server := gin.Default()
	
	log.Print("Migrating models")
	db.AutoMigrate(&notes)

	// db.Create(&notes)

	notes_endpoint := server.Group("/notes")

	log.Print("Setting up note endpoints")
	notes_endpoint.GET("/", getNotes)
	notes_endpoint.POST("/", createNote)

	notes_endpoint.GET("/:id", getNote)
	notes_endpoint.PUT("/:id", updateNote)
	notes_endpoint.DELETE("/:id", deleteNote)

	return server
}

func setupGorm(dialector gorm.Dialector){
	db, err = gorm.Open(dialector,
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

	if err!= nil {
		log.Panic("failed to connect to database: ", err.Error())
	}
}

func main() {
	setupGorm(sqlite.Open("test.db"))

	server := SetupServer()
	
	server.Run(":9090")
}