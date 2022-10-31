package controllers

import (
	core "employe/core/db"
	em "employe/core/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetDataEmploye(c *gin.Context) {

	db := core.GetConnection()
	defer db.Close()

	rows, err := db.QueryContext(c, "SELECT * FROM employes")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	data := []em.Employe{}

	for rows.Next() {
		var each = em.Employe{}
		var err = rows.Scan(&each.Id, &each.Name)
		if err != nil {
			panic(err)
		}
		data = append(data, each)
	}
	c.JSON(200, gin.H{
		"message": "Hello World",
		"data":    data,
	})
}

func GetDataEmployeById(c *gin.Context) {

	db := core.GetConnection()
	defer db.Close()

	id := c.Param("id")
	rows, err := db.QueryContext(c, "SELECT * FROM employes WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	data := []em.Employe{}

	for rows.Next() {
		var each = em.Employe{}
		var err = rows.Scan(&each.Id, &each.Name)
		if err != nil {
			panic(err)
		}
		data = append(data, each)
	}
	if len(data) > 0 {
		c.JSON(200, gin.H{
			"message": "Hello World",
			"data":    data[0],
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Hello World",
			"data":    nil,
		})
	}
}

func AddEmploye(c *gin.Context) {
	id := uuid.New()
	db := core.GetConnection()
	defer db.Close()

	name := c.PostForm("name")
	_, err := db.ExecContext(c, "INSERT INTO employes (id, name) VALUES (?, ?)", id, name)
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"message": "Hello World",
		"data":    id,
	})

}

func DeleteEmploye(c *gin.Context) {
	db := core.GetConnection()
	defer db.Close()

	id := c.Param("id")
	_, err := db.ExecContext(c, "DELETE FROM employes WHERE id = ?", id)
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"message": "Hello World",
		"data":    id,
	})

}
