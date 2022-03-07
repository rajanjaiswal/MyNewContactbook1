package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pagination(c *gin.Context) (int, int, error) {
	pageStr := c.Query("page")
	page := 0 // this is to make page always start from 0
	if pageStr != "" {
		p, err := strconv.Atoi(pageStr) //converting string on line 43
		if err != nil {

			return 0, 0, err
		}
		page = p - 1
		if page <= 0 {
			page = 0
		}

	}

	sizeString := c.Query("size") //this is to make size 20 in one single page
	limit := 20
	if sizeString != "" {
		l, err := strconv.Atoi(sizeString)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return 0, 0, err
		}
		limit = l
		if limit <= 0 {
			limit = 20
		}
	}
	return page, limit, nil
}
