package utils

import db "qlist/db/sqlc"

func GetIsUserOwningEvent(events []db.Event, eventId int) bool {
	for _, event := range events {
		if event.ID == int32(eventId) {
			return true
		}
	}
	return false
}
