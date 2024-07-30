package handlers

import (
	"net/http"
	"todo-list-api/db"
	"todo-list-api/models"

	"github.com/gin-gonic/gin"
)

func GetTask(c *gin.Context) {
	rows, err := db.DB.Query("select id , title , desce , completed , created_at from task")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error ": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.Id, &task.Title, &task.Desce, &task.Completed, &task.Created_at); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err := db.DB.Exec("insert into task (title , desce , completed) values (?,?,?)", task.Title, task.Desce, task.Completed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err := db.DB.Exec("update task set title = ? , desce = ? , completed = ? where id = ?", task.Title, task.Desce, task.Completed, task.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	_, err := db.DB.Exec("delete from task where id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
