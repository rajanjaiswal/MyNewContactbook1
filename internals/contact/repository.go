package contact

import "gorm.io/gorm"

type RepositoryInterface interface {
	Query(Offset, limit int, query string) ([]Contact, error)
	Get(id uint) (Contact, error)
	Create(req *Contact) error
	Update(id uint, update *Contact) error
	Delete(id uint) error
}

type repository struct {
	db gorm.DB
}

func NewRepository(db gorm.DB) RepositoryInterface {
	return &repository{db}

}

func (repository *repository) Query(offset, limit int, query string) ([]Contact, error) {
	var dataList []Contact
	err := repository.db.
		Debug(). //getting data from database to API  //Debug in this lines helps to know the debug query in terminal
		Model(&Contact{}).
		//if filterString != "" {
		Preload("Address").Preload("Phone").Preload("Languages").
		Where("first_name like ? ", "%"+query+"%").
		Limit(limit).
		Offset(offset).
		Find(&dataList).
		Error
	return dataList, err

}
func (repository *repository) Get(id uint) (Contact, error) {
	contact := Contact{}

	err := repository.db.Debug().
		Model(&Contact{}).
		Preload("Address").Preload("Phone").Preload("Languages").
		First(&contact, id).Error
	return contact, err

}
func (repository *repository) Create(req *Contact) error {
	return repository.db.
		Debug().
		Model(&Contact{}).
		Create(&req).
		Error

}
func (repository *repository) Update(id uint, update *Contact) error {
	return repository.db.
		Debug().
		Model(&Contact{}).
		Where("id = ?", id).
		Updates(&update).Error

}
func (repository *repository) Delete(id uint) error {
	return repository.db.
		Debug().
		Delete(&Contact{}, id).
		Error

}
