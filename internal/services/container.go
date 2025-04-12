package services

type Container struct {
	Item *ItemService
}

func NewCotainer() *Container {
	return &Container{Item: NewItemService()}
}
