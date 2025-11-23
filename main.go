package main

import (
	"ticket-printing-system/dto"
	"ticket-printing-system/handler"
	"ticket-printing-system/service"
)

func main() {
	req := dto.NewRequest("Sidik", "Jakarta")
	// Init service and handler
	service := service.NewTicketService()
	handler := handler.NewTicketHandler(service)
	handler.Request(req)
}
