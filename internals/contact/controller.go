package contact

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rajanjaiswal/MyNewContactbook/pkg/utils"
)

type resource struct {
	service ServiceInterface
}

func RegisterRoutes(r *gin.Engine, service ServiceInterface) {
	resource := &resource{service}
	r.GET("/contacts", resource.Query)

	r.POST("/contacts", resource.Create)

	r.PUT("/contacts/:id", resource.Update)
	r.GET("/contacts/:id", resource.Get)
	r.DELETE("/contacts/:id", resource.Delete)

}

// this is for query parameters from line 11 to 39
func (resource *resource) Query(c *gin.Context) {

	//this is query parameters

	// TODO fetch contact list from database where first_name = "ram"
	// select * from contacts where first_name = {first_name} and last_name = "last_name"

	page, limit, err := utils.Pagination(c) // this for pagination and limit from utils till line 27
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	// filterString := "%" + first_name + "%"
	// fmt.Println(filterString)

	contactList, err := resource.service.Query(page*limit, limit, c.Query("first_name"))

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "getting all contacts",
		"data":    contactList,
	})
}

func (resource *resource) Create(c *gin.Context) {
	//	body, _ := ioutil.ReadAll(c.Request.Body)  //THIS IS OLD WAY FROM LINE 36,37,38
	//	fmt.Println(string(body))
	//	err := json.Unmarshal(body, &contact)
	contact := Contact{}
	if err := c.Bind(&contact);
	//err := c.BindJSON(&contact) //THIS IS SAME AS LINE 40 BUT ONLY DO JSON FILE input
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// taking data froM API TO DATABASE
	contact, err := resource.service.Create(&contact)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}
	//contact.ID = 5 // This is for the test

	c.JSON(http.StatusOK, gin.H{
		"message": "Create contacts Sucessfully",
		"data":    contact,
	})

}

func (resource *resource) Update(c *gin.Context) {
	contact := Contact{}
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	_, err := resource.service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Record not found",
		})
		return
	}

	if err := c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Updating contact whose value is onlu changed
	// update := make(map[string]interface{})
	// if contactUpdate.FirstName != "" && contactFromDatabase.FirstName != contactUpdate.FirstName{
	// 	update["first_name"] = contactUpdate.FirstName

	// }
	cont, err := resource.service.Update(uint(id), &contact)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update the contact",
		"data":    cont,
	})

}

func (resource *resource) Delete(c *gin.Context) { //this is to delect by id
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	err := resource.service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Successfully Delected",
	})
}

//this is for path parameters from line 76 to 86
func (resource *resource) Get(c *gin.Context) { // this to filter by using id or Query in database to fetch contact by id
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	// Atoi
	// select * from contacts where id = {id}
	contact, err := resource.service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Contact By Id ",
		"data":    contact,
	})
}
