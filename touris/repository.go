package touris

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Touris, error)
	FindById(ID int) (Touris, error)
	FindByRating(rating float32) ([]Touris, error)
	FindByLokasi(lokasi string) ([]Touris, error)
	FindByAll(rating float32, lokasi string) ([]Touris, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Touris, error) {
	var touris []Touris

	err := r.db.Find(&touris).Error
	if err != nil {
		return touris, err
	}
	return touris, nil
}

func (r *repository) FindById(ID int) (Touris, error) {
	var touris Touris

	err := r.db.Where("id = ?", ID).Find(&touris).Error

	if err != nil {
		return touris, err
	}
	return touris, nil
}

func (r *repository) FindByAll(rating float32, lokasi string) ([]Touris, error) {
	var all []Touris

	err := r.db.Where("rating = ? and location like ?", rating, "%"+lokasi+"%").Find(&all).Error
	if err != nil {
		return all, err
	}
	return all, nil

}

func (r *repository) FindByRating(rating float32) ([]Touris, error) {
	//inisiasi user yang bertujuan dalam mencari email yang ada di struct User
	var rate []Touris

	//menggunakan where yang mana email mana ni yang mau di loginin sesuai ga
	err := r.db.Where("rating = ?", rating).Find(&rate).Error
	if err != nil {
		return rate, err
	}
	return rate, nil
}

func (r *repository) FindByLokasi(rating string) ([]Touris, error) {
	var lokasi []Touris

	err := r.db.Where("location like ?", "%"+rating+"%").Find(&lokasi).Error
	if err != nil {
		return lokasi, err
	}
	return lokasi, nil
}
