package product

type Service struct {
}

func NewService() {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProduct
}
