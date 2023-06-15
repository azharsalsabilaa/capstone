package touris

type Service interface {
	GetTouris(ID int) ([]Touris, error)
	GetRating(input InputGetTouris) ([]Touris, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTouris(ID int) ([]Touris, error) {
	cafe, err := s.repository.FindAll()
	if err != nil {
		return cafe, err
	}
	return cafe, nil
}

func (s *service) GetRating(input InputGetTouris) ([]Touris, error) {
	if input.Rating != 0 && input.Lokasi != "" {
		newTouris, err := s.repository.FindByAll(input.Rating, input.Lokasi)
		if err != nil {
			return newTouris, err
		}
		return newTouris, nil
	} else if input.Lokasi != "" {
		newLokasi, err := s.repository.FindByLokasi(input.Lokasi)
		if err != nil {
			return newLokasi, err
		}
		return newLokasi, nil
	} else {
		newRating, err := s.repository.FindByRating(input.Rating)
		if err != nil {
			return newRating, err
		}
		return newRating, nil
	}
}
