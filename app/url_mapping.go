package app

import (
	"github.com/sachin120/contact_app/controllers/contacts"
	"github.com/sachin120/contact_app/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/contacts/:contact_id", contacts.GetContact)
	router.POST("/contacts", contacts.CreateContact)
	router.DELETE("/contacts/:contact_id", contacts.DeleteContact)
}
