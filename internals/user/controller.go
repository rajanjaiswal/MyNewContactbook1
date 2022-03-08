package user

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
	r.GET("/users", resource.Query)

	r.POST("/users", resource.Create)

	r.PUT("/users/:id", resource.Update)
	r.GET("/users/:id", resource.Get)
	r.DELETE("/users/:id", resource.Delete)

}

// this is for query parameters from line 11 to 39
func (resource *resource) Query(c *gin.Context) {

	//this is query parameters

	// TODO fetch User list from database where first_name = "ram"
	// select * from Users where first_name = {first_name} and last_name = "last_name"

	page, limit, err := utils.Pagination(c) // this for pagination and limit from utils till line 27
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	// filterString := "%" + first_name + "%"
	// fmt.Println(filterString)

	userList, err := resource.service.Query(page*limit, limit, c.Query("first_name"))

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "getting all users",
		"data":    userList,
	})
}

func (resource *resource) Create(c *gin.Context) {
	//	body, _ := ioutil.ReadAll(c.Request.Body)  //THIS IS OLD WAY FROM LINE 36,37,38
	//	fmt.Println(string(body))
	//	err := json.Unmarshal(body, &User)
	user := User{}
	if err := c.Bind(&user);
	//err := c.BindJSON(&User) //THIS IS SAME AS LINE 40 BUT ONLY DO JSON FILE input
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// taking data froM API TO DATABASE
	user, err := resource.service.Create(&user)

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}
	//User.ID = 5 // This is for the test

	c.JSON(http.StatusOK, gin.H{
		"message": "Create users Sucessfully",
		"data":    user,
	})

}

func (resource *resource) Update(c *gin.Context) {
	user := User{}
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	_, err := resource.service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Record not found",
		})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Updating User whose value is onlu changed
	// update := make(map[string]interface{})
	// if UserUpdate.FirstName != "" && UserFromDatabase.FirstName != UserUpdate.FirstName{
	// 	update["first_name"] = UserUpdate.FirstName

	// }
	cont, err := resource.service.Update(uint(id), &user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update the user",
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
func (resource *resource) Get(c *gin.Context) { // this to filter by using id or Query in database to fetch User by id
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	// Atoi
	// select * from Users where id = {id}
	user, err := resource.service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get User By Id ",
		"data":    user,
	})
}
