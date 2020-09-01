package contacts

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/sachin120/contact_app/services"
	"github.com/sachin120/contact_app/utils/errors"
	"github.com/sachin120/contact_app/utils/random"

	"github.com/sachin120/contact_app/domain/contacts"

	"github.com/gin-gonic/gin"
)

// GetContact ....
func GetContact(c *gin.Context) {
	var cnt contacts.Contact
	contactid := c.Param("contact_id")
	i, err := strconv.Atoi(contactid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "invalid id")
		return
	}
	cnt.ID = i
	cnt2, serr := services.FindContact(cnt)
	if serr != nil {
		log.Println(serr)
		restErr := errors.NewBadRequestError("Result not found.") // Change the error from bad request to something resonable
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, &cnt2)
}

// CreateContact ...
func CreateContact(c *gin.Context) {
	var cnt contacts.Contact
	if err := c.ShouldBindJSON(&cnt); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	cnt.ID = random.GetRandom()
	result, err := services.CreateContact(cnt)
	if err != nil {
		fmt.Print(result)
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// DeleteContact ...
func DeleteContact(c *gin.Context) {
	var cnt contacts.Contact
	contactid := c.Param("contact_id")
	i, err := strconv.Atoi(contactid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "invalid id")
		return
	}
	cnt.ID = i
	serr := services.DeleteContact(cnt)
	if serr != nil {
		log.Println(serr)
		restErr := errors.NewBadRequestError("Result not found.") // Change the error from bad request to something resonable
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, "Deleted")
}
