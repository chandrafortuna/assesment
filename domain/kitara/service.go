package kitara

type Service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return Service{
		repo: r,
	}
}

func (s *Service) InitData(variant ProductVariant) (*ProductVariant, error) {
	res, err := s.repo.Save(&variant)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) GetAll() ([]*ProductVariant, error) {
	variants, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return variants, nil
}

func (s *Service) Reserve(req ReserveProductReq) (*ReserveProductResponse, error) {
	variant, err := s.repo.GetByID(req.ProductVariantID)
	if err != nil {
		return NewReserveProductResponse(false, err.Error(), &req), nil
	}

	err = variant.ReduceQty(req.Qty)
	if err != nil {
		return NewReserveProductResponse(false, err.Error(), &req), nil
	}

	_, err = s.repo.Save(variant)
	return NewReserveProductResponse(true, "", &req), nil
}

func (s *Service) Clear() (bool, error) {
	err := s.repo.Clear()
	if err != nil {
		return false, err
	}

	return true, nil
}
