package module1

// Budget stores budget information
type budget struct {
	Max   float32
	Items []Item
}

// Item stores item information
type item struct {
	Description string
	Price       float32
}
