package main

import (
	"net/http"
  
	"github.com/gin-gonic/gin"
  )
  
import "fmt"

func getAllBooks(c *gin.Context){
c.IndentedJSON(http.StatusOK,books)
}

func postBooks(c *gin.Context){
	var newBook book
	c.BindJSON(&newBook)
	books = append(books,newBook)
	c.IndentedJSON(http.StatusCreated,newBook)
}

func getABook(c *gin.Context){
var id = c.Param("id")

for _,book :=range books{
	if book.Id == id{  
    c.IndentedJSON(http.StatusOK,book)
	return
	}
}
	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Book not Found"})
}

type book struct{
	Id string `json:"id"`
	Title string 
	Author string 
	Price string 
}

var books = []book{
	{Id:"1",Title:"Half girlfriend",Author:"chetan jii",Price:"145000"},
	{Id:"2",Title:"Inferno",Author:"Da Vinci",Price:"9000"},
	{Id:"3",Title:"Oxford",Author:"British",Price:"50"},
}
func main(){
	fmt.Printf("dcds")

	router := gin.Default()
	router.GET("/books",getAllBooks)
	router.POST("/books",postBooks)
	router.GET("/books/:id",getABook)
	router.Run("localhost:8080")
}