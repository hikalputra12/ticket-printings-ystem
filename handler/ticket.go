package handler

//bertuga smenampilkan informasi tiket ke konsol

import (
	"fmt"
	"ticket-printing-system/dto"
	"ticket-printing-system/service"
)

type TicketHandler struct {
	Ticketservice service.TicketService
}

// constructor
func NewTicketHandler(ticketservice service.TicketService) TicketHandler {
	return TicketHandler{
		Ticketservice: ticketservice,
	}
}

/*
NewRequestByUser handle permintaan tiket baru dari user
menerima input dari user dan mengembalikan response yang berisi informasi tiket
bisa juga mengembalikan error jika ada masalah dalam proses pembuatan request
fungsi ini akan mencetak informasi tiket ke konsol
*/
func (tickethandler *TicketHandler) Request(req dto.Request) {

	response, err := tickethandler.Ticketservice.NewRequestByUser(req)
	if err != nil {
		fmt.Printf("error creating new request: %v", err)
	} else {
		fmt.Println("===Harga Tiket ===")
		fmt.Printf("Penumpang: %s\n", response.Name)
		fmt.Printf("Tujuan: %s\n", response.Destination)
		fmt.Printf("Harga: Rp  %.2f\n", float64(response.Price))
		fmt.Println("===================")
	}
}
