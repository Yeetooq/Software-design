package models

type Order struct {
	ID        int
	ClientID  int
	PartID    int
	Quantity  int
	TotalCost float64
}
