package main

import (
	"fmt"
	"ticket-printing-system/dto"
	"ticket-printing-system/handler"
	"ticket-printing-system/service"
)

func main() {
	var name string
	var destination string
	// Init service and handler
	service := service.NewTicketService()
	handler := handler.NewTicketHandler(service)
	fmt.Print("masukkan nama anda: ")
	fmt.Scanln(&name)
	fmt.Print("masukkan destinasi tujuan anda: ")
	fmt.Scanln(&destination)
	//membuat request baru
	req := dto.NewRequest(name, destination) //sebanarnya bisa langsung namun saya mencoba meminta user memasukkan data melalui CLI
	//memanggil fungsi request pada handler
	handler.Request(req)

}
