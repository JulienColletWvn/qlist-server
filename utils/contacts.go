package utils

import db "qlist/db/sqlc"

func GetIsUserOwningContact(contacts []db.Contact, contactId int) bool {
	for _, contact := range contacts {
		if contact.ID == int32(contactId) {
			return true
		}
	}
	return false
}
