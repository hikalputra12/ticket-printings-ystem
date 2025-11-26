package utils

// file.go berisi fungsi untuk membaca dan menulis file JSON
// atau sebagai pembantu
import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath" // Tambahkan import filepath jika Anda ingin menggunakan os.MkdirAll("data") dengan lebih robust
	"ticket-printing-system/model"
)

// file paths berisi lokasi dan harga
const DestinationsFilePath = "data/destination.json"

// fungsi untuk mengecek file ada atau tidak dan jika file kosong akan menghasilkan array kodong
func EnsureUsersFile() error {
	_, err := os.Stat(DestinationsFilePath)

	// filepath.Dir() untuk membuat direktori jika belum ada
	//Menentukan apakah error yang diberikan (err) setara secara semantik dengan error target (target)
	if errors.Is(err, os.ErrNotExist) { //cek err apakah sama atau tidak
		if err := os.MkdirAll(filepath.Dir(DestinationsFilePath), 0755); err != nil { //hasil error maka buat file baru

			return err
		}
		return os.WriteFile(DestinationsFilePath, []byte("[]"), 0644) //menulis file nya
	}
	return nil //jika tidak ada error maka return nil
}

// fungsi untuk membaca destinasi dari file
func ReadDestinationsFromFile() ([]model.Ticket, error) { // Tipe kembalian menggunakan []model.Tiket
	//cek file ada atau tidak
	if err := EnsureUsersFile(); err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(DestinationsFilePath)
	if err != nil {
		return nil, err
	}

	var routes []model.Ticket // Menggunakan model.Ticket secara konsisten
	if err := json.Unmarshal(bytes, &routes); err != nil {
		// Tampilkan error unmarshal yang lebih deskriptif
		return nil, errors.New("gagal mengurai data destinasi: " + err.Error())
	}

	return routes, nil
}
