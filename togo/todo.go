package main 

import (
	"github.com/gin-gonic/gin" 
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *grom.DB
func init() {
	// open a db connection
	var err error		// db driver: mysql,   userName: root,  password: 12345,  dbName: demo
	db, err =grom.Open("mysql", "root:12345@/demo?  
	charset=utf8&parseTime=True&loc-Local")
	if err != nil {
		panic("fauled to connet database")
	}
	//Migrate the schema
	db.AutoMigrate(&todoModel{})
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", createTodo)
		v1.GET("/", fetchAllTodo)
		v1.GET("/:id", fetchSingleTodo)
		v1.PUT("/:id", updateTodo)
		v1.DELETE("/:id", deleteTodo)
	}	
	router.Run()
}

type (
	// todoModel describes a todoModel type
	todoModel struct {
		gorm.Model  // ID, CreateAt, UpdateAt, DeleteAt를 포함하는 Model 구조체
		Title		string 'json:"title"'
		Completed 	int 	'json:"completed"'
	}

	// transformedTodo represents a formatted todo
	transformedTodo struct {
		ID 			uint 	'json:"id"'
		Title 		string 	'json:"title"'
		COmpleted 	bool 	'json:"completed'
	}
)