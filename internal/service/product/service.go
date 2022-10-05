package product

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(idx int) (*Product, error) {
	return &allProducts[idx], nil
}

func (s *Service) Delete(idx int) ([]Product, error) {
	allProducts = append(allProducts[:idx], allProducts[idx+1:]...)
	return allProducts, nil
}

func (s *Service) Create(Entity string) ([]Product, error) {
	allProducts = append(allProducts, Product{Title: Entity})
	return allProducts, nil
}

func (s *Service) Replace(idx int,newArg string) ([]Product, error) {
	rightProducts := allProducts[idx+1:]
	leftProducts := append(allProducts[:idx], Product{Title:newArg})
	allProducts = append(leftProducts, rightProducts...)
	return allProducts, nil
}
