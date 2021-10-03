package user

import (
	"log"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]User, error)
	GetById(uint64) (User, error)
	GetIn([]uint64) ([]User, error)
	Create(User) (User, error)
	Update(User) (User, error)
	Delete(uint64) error
	Migrate() error
}

type sqliteUserRepository struct {
	Db *gorm.DB
}

func NewSqliteRepository(db *gorm.DB) UserRepository {
	repo := sqliteUserRepository{
		Db: db,
	}

	err := repo.Migrate()
	if err != nil {
		log.Fatal("Repo migration failed", err)
	}

	return repo
}

func (s sqliteUserRepository) GetAll() (users []User, err error) {
	log.Println("{SQLITE USER REPOSITORY} GetAll")

	err = s.Db.Find(&users).Error

	return users, err
}

func (s sqliteUserRepository) GetById(userId uint64) (user User, err error) {
	log.Println("{SQLITE USER REPOSITORY} GetById")

	err = s.Db.First(&user, userId).Error

	return user, err
}

func (s sqliteUserRepository) GetIn(userIds []uint64) (users []User, err error) {
	log.Println("{SQLITE USER REPOSITORY} GetIn")

	err = s.Db.Where(userIds).Find(&users).Error

	return users, err
}

func (s sqliteUserRepository) Create(user User) (User, error) {
	log.Println("{SQLITE USER REPOSITORY} Create")

	err := s.Db.Create(&user).Error

	return user, err
}

func (s sqliteUserRepository) Update(user User) (User, error) {
	log.Println("{SQLITE USER REPOSITORY} Update")

	err := s.Db.Save(user).Error

	return user, err
}

func (s sqliteUserRepository) Delete(id uint64) error {
	log.Println("{SQLITE USER REPOSITORY} Delete")

	err := s.Db.Delete(&User{}, id).Error

	return err
}

func (s sqliteUserRepository) Migrate() error {
	log.Println("{SQLITE USER REPOSITORY} Create")

	err := s.Db.AutoMigrate(&User{})

	return err
}
