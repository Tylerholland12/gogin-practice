package newsfeed

type Getter interface {
	GetAll() []Item
}

// function parameters
type Adder interface {
	Add() (item Item)
}

type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Repo struct {
	Items []Item
}

func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

// append items to slice every time add is called
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	// return slice of items
	return r.Items
}
