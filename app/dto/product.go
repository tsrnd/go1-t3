package dto

type DataTransferProduct struct {
	Id           uint
	Name         string
	Descrition   string
	Quantity     int
	Price        int
	PrimaryImage string
	Images       map[int][]string
}
