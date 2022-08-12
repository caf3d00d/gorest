package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// A struct defines animal's properties
type animal struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Sex  string `json:"sex"`
}

var animals = []animal{
	{ID: "1", Name: "Blinka", Type: "Dog", Sex: "Female"},
	{ID: "2", Name: "Karma", Type: "Cat", Sex: "Female"},
	{ID: "3", Name: "Row", Type: "Dog", Sex: "Male"},
}

/* @getAnimals
A function to get a list of animals
curl http://localhost:8080/animals
*/
func getAnimals(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, animals)
}

/* @getAnimalByID
A function to get an animal by id
curl http://localhost:8080/animals/2
*/
func getAnimalByID(context *gin.Context) {
	id := context.Param("id")

	for _, animal := range animals {
		if animal.ID == id {
			context.IndentedJSON(http.StatusOK, animal)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Animal not found"})
}

/* @addAnimal
A function that adds one animal to the list
curl http://localhost:8080/addanimal --include --header "Content-Type: application/json" \
	--request "POST" \
	--data '{"id": "4", "Name": "Col", "Type": "Cat", "Sex": "Male"}'
*/
func addAnimal(context *gin.Context) {
	var newAnimal animal

	err := context.BindJSON(&newAnimal)
	if err != nil {
		return
	}

	animals = append(animals, newAnimal)
	context.IndentedJSON(http.StatusCreated, newAnimal)
}

func main() {
	router := gin.Default()
	router.GET("/animals", getAnimals)
	router.GET("/animals/:id", getAnimalByID)
	router.POST("/addanimal", addAnimal)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
