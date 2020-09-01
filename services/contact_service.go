package services

import (
	"github.com/sachin120/contact_app/domain/contacts"
	"github.com/sachin120/contact_app/utils/errors"
)

// CreateContact ...
func CreateContact(c contacts.Contact) (*contacts.Contact, *errors.RestErr) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	if err := c.Save(); err != nil {
		return nil, err
	}
	return &c, nil
}

// FindContact ...
func FindContact(c contacts.Contact) (*contacts.Contact, *errors.RestErr) {
	if err := c.Get(); err != nil {
		return nil, err
	}
	return &c, nil
}

// DeleteContact ...
func DeleteContact(c contacts.Contact) *errors.RestErr {
	return c.Delete()
}
