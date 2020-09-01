package contacts

import (
	"github.com/sachin120/contact_app/utils/errors"
	"github.com/sachin120/contact_app/utils/validate"
)

// Contact ...
type Contact struct {
	ID        int    `json:"ID"`
	FirstName string `json:"first_name"`
	LastnName string `json:"last_name"`
	Prinumber string `json:"prinumber"`
	Address   string `json:"address"`
	Emailid   string `json:"emailid"`
}

// Validate ...
func (c *Contact) Validate() *errors.RestErr {
	if !validate.IsEmailValid(c.Emailid) {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
