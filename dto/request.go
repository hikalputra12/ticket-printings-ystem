package dto

type Request struct {
	Name, Destination string
}

// membuat request baru
func NewRequest(name, destination string) Request {
	return Request{
		Name:        name,
		Destination: destination,
	}
}
