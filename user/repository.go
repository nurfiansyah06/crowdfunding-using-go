package user

import "gorm.io/gorm"

// create interface for function
type Repository interface {
	Save(user User) (User, error)
	FindbyEmail(email string) (User, error)
	FindbyID(ID int) (User, error)
	Update(user User) (User, error)
}

// create repository struct for connect to db use gorm
type repository struct {
	db *gorm.DB
}


func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}


// func create user to db
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindbyEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindbyID(ID int) (User, error) {
	var user User
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error){
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}