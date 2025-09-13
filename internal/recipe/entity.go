package recipe

type ItemDetail struct {
	ItemID   string
	Quantity int
}

type Recipe struct {
	ID      string
	GroupID string
	Name    string
	Inputs  []ItemDetail
	Outputs []ItemDetail
}
