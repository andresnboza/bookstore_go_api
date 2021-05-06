package services

var (
	ItemsService itemsServiceinterface = &itemsService{}
)

type itemsServiceinterface interface {
	GetItem()
	SaveItem()
}

type itemsService struct {
}

func (s *itemsService) GetItem() {

}

func (s *itemsService) SaveItem() {

}
