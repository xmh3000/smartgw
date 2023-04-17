package repository

import (
	"encoding/json"
	"errors"
	"github.com/boltdb/bolt"
	"smartgw/api/domain"
	"smartgw/lib/database"
)

type UserRepository interface {
	Save(user *domain.User) error
	Delete(username string) error
	Find(username string) (domain.User, error)
	FindAll() ([]domain.User, error)
	Migrate() error
}

var _ UserRepository = (*userRepository)(nil)

type userRepository struct {
	db *bolt.DB
}

func NewUserRepository(db *bolt.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Save(user *domain.User) error {
	return u.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.User))
		if data, err := json.Marshal(user); err != nil {
			return err
		} else {
			return b.Put([]byte(user.Username), data)
		}
	})
}

func (u *userRepository) Delete(username string) error {
	return u.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.User))

		return b.Delete([]byte(username))
	})
}

func (u *userRepository) Find(username string) (domain.User, error) {
	user := domain.User{}
	err := u.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.User))
		data := b.Get([]byte(username))

		if data != nil {
			return json.Unmarshal(data, &user)
		} else {
			return errors.New("没有找到数据")
		}
	})

	return user, err
}

func (u *userRepository) FindAll() ([]domain.User, error) {
	users := make([]domain.User, 0)

	err := u.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(database.User))
		return b.ForEach(func(k, v []byte) error {
			user := domain.User{}
			err := json.Unmarshal(v, &user)
			if err == nil {
				users = append(users, user)
			}
			return err
		})
	})

	return users, err
}

func (u *userRepository) Migrate() error {
	if _, err := u.Find("admin"); err != nil {
		return u.Save(&domain.User{
			Username: "admin",
			Password: "123456",
		})
	}

	return nil
}
