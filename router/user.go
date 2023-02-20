package router

import (
	user "qlist/handler/user"
	userEvents "qlist/handler/user/events"

	"github.com/gofiber/fiber/v2"
)

func User(r fiber.Router) {
	r.Get("/", user.GetUser)

	contacts := r.Group("/contacts")
	contacts.Get("/", user.GetUserContacts)
	contacts.Get("/:contactId", user.GetUserContact)
	contacts.Post("/", user.CreateContacts)
	contacts.Delete("/:contactId", user.DeleteUserContact)

	events := r.Group("/events")
	events.Get("/", userEvents.GetEvents)
	events.Get("/:eventId", userEvents.GetEvent)
	events.Post("/", userEvents.CreateEvent)
	events.Delete("/:eventId", userEvents.DeleteUserEvent)

	guests := events.Group("/:eventId/guests")
	guests.Get("/", userEvents.GetUserEventGuests)
	guests.Get("/:guestId", userEvents.GetUserEventGuest)
	guests.Post("/", userEvents.CreateUserEventGuests)
	guests.Delete("/:guestId", userEvents.DeleteUserEventGuest)

	ticketsTypes := events.Group("/:eventId/ticketsTypes")
	ticketsTypes.Get("/", userEvents.GetUserEventTicketTypes)
	ticketsTypes.Get("/:ticketTypeId", userEvents.GetUserEventTicketType)
	ticketsTypes.Post("/", userEvents.CreateUserEventTicketType)
	ticketsTypes.Delete("/:ticketTypeId", userEvents.DeleteUserEventTicketType)

	tickets := guests.Group("/:guestId/tickets")
	tickets.Get("/", userEvents.GetGuestTickets)
	tickets.Get("/:ticketId", userEvents.GetGuestTickets)
	tickets.Post("/", userEvents.CreateGuestTicket)
	tickets.Delete("/:ticketId", userEvents.DeleteGuestTicket)

	ticketTransactions := tickets.Group("/:ticketId/transactions")
	ticketTransactions.Get("/", userEvents.GetGuestTicketTransactions)
	ticketTransactions.Get("/:transactionId", userEvents.GetGuestTicketTransaction)
	ticketTransactions.Post("/", userEvents.CreateGuestTicketTransaction)
	ticketTransactions.Put("/:transactionId/status", userEvents.UpdateGuestTicketTransactionStatus)

	cashiers := events.Group("/:eventId/cashiers")
	cashiers.Get("/", userEvents.GetUserEventCashiers)
	cashiers.Get("/:cashierId", userEvents.GetUserEventCashier)
	cashiers.Post("/", userEvents.CreateUserEventCashier)
	cashiers.Delete("/:cashierId", userEvents.DeleteUserEventCashier)

	sellers := events.Group("/:eventId/sellers")
	sellers.Get("/", userEvents.GetUserEventSellers)
	sellers.Get("/:sellerId", userEvents.GetUserEventSeller)
	sellers.Post("/", userEvents.CreateUserEventSeller)
	sellers.Delete("/:sellerId", userEvents.DeleteUserEventSeller)

	stewards := events.Group("/:eventId/stewards")
	stewards.Get("/", userEvents.GetUserEventStewards)
	stewards.Get("/:stewardId", userEvents.GetUserEventSteward)
	stewards.Post("/", userEvents.CreateUserEventSteward)
	stewards.Delete("/:stewardId", userEvents.DeleteUserEventSteward)

	walletsTypes := events.Group("/:eventId/walletsTypes")
	walletsTypes.Get("/", userEvents.GetUserEventWalletTypes)
	walletsTypes.Get("/:walletTypeId", userEvents.GetUserEventWalletType)
	walletsTypes.Post("/", userEvents.CreateUserEventWalletType)
	walletsTypes.Delete("/:walletTypeId", userEvents.DeleteUserEventWalletType)

	wallets := events.Group(":guestId/wallets")
	wallets.Get("/", userEvents.GetGuestWallets)
	wallets.Get("/:walletId", userEvents.GetGuestWallet)
	wallets.Post("/", userEvents.CreateGuestWallet)
	wallets.Delete("/:walletId", userEvents.DeleteGuestWallet)

	statistics := events.Group(("/:eventId/statistics"))
	statistics.Get("/", userEvents.GetUserEventStatistics)
}
