package utils

// file.go berisi fungsi untuk membaca dan menulis file JSON (destinations)
// sebagai pembantu untuk service
import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"ticket-printing-system/model"
)

// DestinationsFilePath adalah lokasi file destinasi
const DestinationsFilePath = "data/destinations.json"

// EnsureDestinationsFile memastikan file destinasi ada, jika tidak ada dibuatkan
func EnsureDestinationsFile() error {
	_, err := os.Stat(DestinationsFilePath)
	if errors.Is(err, os.ErrNotExist) {
		// buat direktori jika belum ada
		if err := os.MkdirAll(filepath.Dir(DestinationsFilePath), 0o755); err != nil {
			return err
		}
		return os.WriteFile(DestinationsFilePath, []byte("[]"), 0o644)
	}
	return nil
}

// ReadDestinationsFromFile membaca destinasi dari file JSON dan mengembalikan slice model.Ticket
func ReadDestinationsFromFile() ([]model.Ticket, error) {
	if err := EnsureDestinationsFile(); err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(DestinationsFilePath)
	if err != nil {
		return nil, err
	}

	var tickets []model.Ticket
	if err := json.Unmarshal(bytes, &tickets); err != nil {
		return nil, fmt.Errorf("gagal mengurai data destinasi: %w", err)
	}

	return tickets, nil
}

// WriteDestinationsToFile menulis slice model.Ticket ke file destinasi
func WriteDestinationsToFile(tickets []model.Ticket) error {
	bytes, err := json.MarshalIndent(tickets, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(DestinationsFilePath, bytes, 0o644)
}

// backward-compatible alias (opsional) jika ada kode lama yang menggunakan EnsureUsersFile
func EnsureUsersFile() error {
	return EnsureDestinationsFile()
}
