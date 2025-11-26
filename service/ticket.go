package service

import (
	"errors"
	"strings"
	"ticket-printing-system/dto"
	"ticket-printing-system/utils"
)

type TicketService struct{}

// constructor
func NewTicketService() TicketService {
	return TicketService{}
}


// fungsi untuk membuat request baru oleh user
// menerima input dari user dan mengembalikan model.User
// bisa juga mengembalikan error jika ada masalah dalam proses pembuatan request
func (ticketService *TicketService) NewRequestByUser(req dto.Request) (dto.Response, error) {
	//logis bisnis untuk membuat request baru oleh user

	// Validasi sederhana
	if strings.TrimSpace(req.Name) == "" ||
		strings.TrimSpace(req.Destination) == "" {
		return dto.Response{}, errors.New("nama dan destinasi di wajibkan")
	}
	// Baca data existing
	destinationData, err := utils.ReadDestinationsFromFile()
	if err != nil {
		return dto.Response{}, err
	}

	// cek harga tiket
	destinationInput := strings.ToLower(req.Destination)
	for _, destination := range destinationData {
		if strings.ToLower(destination.Destination) == destinationInput {
			// Buat user response baru
			newUserResponse := dto.Response{
				Name:        req.Name,
				Destination: destination.Destination,
				Price:       destination.Price,
			}
			return newUserResponse, nil
		}
	}
	return dto.Response{}, errors.New("maaf destinasi tidak ditemukan")

}
