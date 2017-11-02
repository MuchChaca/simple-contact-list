package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/MuchChaca/Training/simple-contact-list/src/models"
	"github.com/labstack/echo"
)

// H is a small trick allowing to return arbitrary JSON in our response
type H map[string]interface{}

// FindAll endpoint
func FindAll(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Fetch contaacts using the new model
		return c.JSON(http.StatusOK, models.FindAll(db))
	}
}

// AddContact endpoint
func AddContact(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new contact
		contact := models.Contact{}
		// Map incoming JSON body to the new contact
		// TODO: Replace bind with more efficient method
		c.Bind(&contact)
		// Add contact using the model
		id, err := models.AddContact(db, contact.LName, contact.FName, contact.Email, contact.Portable, contact.Phone)
		// Return a JSON response if successful, otherwise an error
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, H{
			"created": id,
		})
	}
}

// ModContact endpoint
func ModContact(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new contact
		contact := models.Contact{}

		// Map incoming JSON body to the new contact
		// TODO: Replace bind with a more efficient method
		c.Bind(&contact)
		// Modify the contact using the model
		_, err := models.ModContact(db, contact)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, H{
			"modified": contact.ID,
		})
	}
}

// DelContact endpoint
func DelContact(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// Use the new model to delete a contact
		_, err := models.DelContact(db, id)
		// Return a JSON response if successful
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, H{
			"deleted": id,
		})
	}
}
