package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rajanjaiswal/MyNewContactbook/models"
)

// this is for query parameters from line 11 to 39
func GetAllContacts(c *gin.Context) {
	first_name := c.Query("first_name")
	last_name := c.Query("last_name") //this is query parameters
	page := strconv.Atoi(c.Query("page"))
	if err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	limit := strconv.Atoi(c.Query("size"))
	if err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}


	// TODO fetch contact list from database where first_name = "ram"
	// select * from contacts where first_name = {first_name} and last_name = "last_name"
	fmt.Println(first_name, last_name)
	var contactList []models.Contact

	filterString := "%" + first_name + "%"
	fmt.Println(filterString)
	query := models.DB.Debug().Model(&models.Contact{})     //getting data from database to API  //Debug in this lines helps to know the debug query in terminal
	if filterString != "" {
		query.where("first_name like ?", filterString)			//this to apply filter for the first name
	}
	
	query.Limit(limit).Offset(limit*page)  //this for paging and limit as line 17 and 23
	err = query.Find(&models.Contact{}).Error 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	c.JSON(http.StatusOK, gin.H{
		"message": "getting all contacts",
		"data":   contactList, //[]models.Contact{
			// {
			// 	ID:        1,
			// 	FirstName: "Ram",
			// 	LastName:  "Sharma",
			// 	Phone:     []models.Phone{},
			// },
			// {
			// 	ID:        2,
			// 	FirstName: "Hari",
			// 	LastName:  "Sharma",
			// 	Phone: []models.Phone{
			// 		{
			// 			PhoneType:   "Mobile",
			// 			PhoneNumber: "9090909090",
			// 		},
			// 	},
			// },
		//},
	})
}

func CreateContact(c *gin.Context) {
	//	body, _ := ioutil.ReadAll(c.Request.Body)  //THIS IS OLD WAY FROM LINE 36,37,38
	//	fmt.Println(string(body))
	//	err := json.Unmarshal(body, &contact)
	contact := models.Contact{}
	err := c.Bind(&contact)
	//err := c.BindJSON(&contact) //THIS IS SAME AS LINE 40 BUT ONLY DO JSON FILE input
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
// taking data froM API TO DATABASE
	err := models.DB.Model(&models.Contact{}).Create(&contact).Error

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}
	contact.ID = 5 // This is for the test

	c.JSON(http.StatusOK, gin.H{
		"message": "Create contacts Sucessfully",
		"data":    contact,
	})

}

func UpdateContacts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update contact comming soon",
	})
}
func DeleteContacts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete contact comming soon",
	})
}

//this is for path parameters from line 76 to 86
func GetContactById(c *gin.Context) {
	id := c.Params.ByName("id")
	// Atoi
	// Query in database to fetch contact by id
	// select * from contacts where id = {id}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Contact By Id ",
		"data":    id,
	})
}