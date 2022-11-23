package handler

// type GuestParams struct {
// 	ContactID int `json:"contactId"`
// }

// func CreateGuest(c *fiber.Ctx) error {
// 	params := GuestParams{}
// 	user := entities.User{}
// 	event := entities.Event{}
// 	contact := entities.Contact{}

// 	eventId := c.AllParams()["eventId"]

// 	userId, err := jwtauth.GetCurrentUserId(c)

// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(err)
// 	}

// 	if err := c.BodyParser(&params); err != nil {
// 		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
// 	}

// 	if eventId == "" {
// 		return c.SendStatus(fiber.StatusBadRequest)
// 	}

// 	utils.Database.Where("id=?", userId).First(&user)
// 	utils.Database.Model(&user).Association("Events").Find(&event)

// 	utils.Database.Model(&user).Where("id=?", params.ContactID).Association("Contacts").Find(&contact)

// 	guest := entities.Guest{
// 		ContactID: contact.ID,
// 	}

// 	e := utils.Database.Model(&event).Association("Guests").Append(&guest)

// 	if e != nil {
// 		return c.SendStatus(fiber.StatusBadRequest)
// 	}

// 	return c.Status(fiber.StatusOK).JSON(&guest)

// }
