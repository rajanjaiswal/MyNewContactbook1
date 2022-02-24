package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllContacts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "getting all contacts",
		"data": []model.contacts{}
	})
}

func GetContacts(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(Body))

	contact := mpdels.Contact{}
	err:=json.Unmarshal(body, &contact)
	if err != nil{
		
	}
		"message": "creat contacts successfuly",
	})

}

func UpdateContacts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "creat contacts comming soon",
	})

}
func DelectContacts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "getting all contacts",
	})

}
