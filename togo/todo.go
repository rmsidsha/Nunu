package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func init() {
	// gorm은 마이그레이션 기능이 있어서 init()을 사용할수 있다.
	// 프로그램을 실행하면 연결을 만든다음 이그레이션을 수행한다
	// open a db connection
	var err error // db driver: mysql,   userName: root,  password: 12345,  dbName: demo
	db, err = gorm.Open("mysql", "todo:123123@/todo?charset=utf8&parseTime=True&loc-Local")
	if err != nil {
		panic("failed to connet database")
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
	// todoModel descrbes a todoModel type
	todoModel struct {
		gorm.Model        // ID, CreateA, UpdateAt, DeleteAt를 포함하는 Model 구조체
		Title      string `json:"title"`
		Completed  int    `json:"completed"`
	}

	// transformedTodo repreents a formatted todo
	transformedTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed`
	}
)

func createTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := todoModel{Title: c.PostForm("title"), Completed: completed}

	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated,
		"message": "Todo item created successfully!", "resourceId": todo.ID})
}

// fetchAllTodo fetch all todos
func fetchAllTodo(c *gin.Context) {
	var todos []todoModel
	var _todos []transformedTodo

	db.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound,
			"messge": "No todo found"})
		return
	}

	// transforms the todos for uilding a good response
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, transformedTodo{ID: item.ID, Title: item.Title,
			Completed: completed})
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})

	}
}
func fetchSingleTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound,
			"messge": "No todo Found"})
		return
	}

	completed := false
	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}

	_todo := transformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
}
func updateTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Status": http.StatusNotFound,
			"messge": "No todo found"})
		return
	}

	db.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("complted"))
	db.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK,
		"message": "Todo updated successfully"})
}
func deleteTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound,
			"messge": "No todo Found"})
		return
	}

	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK,
		"message": "Todo deleted successfully"})
}
