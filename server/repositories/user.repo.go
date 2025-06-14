package repositories

import (
	"database/sql"
	"fmt"
	db "forum/database"
	q "forum/database/query"
	"forum/models"
)

type UserRepository struct {
	BaseRepo
}

func (r *UserRepository) init() {
	r.DB = db.DB
	r.TableName = db.USERS_TABLE
}

func (r *UserRepository) GetUserById(userId string) (models.User, error) {
	row, err := r.DB.GetOneFrom(r.TableName, q.WhereOption{"user_id": userId})
	if err == sql.ErrNoRows {
		return models.User{}, err
	}
	var user models.User
	err = row.Scan(&user.UserId, &user.Firstname, &user.Lastname, &user.Username, &user.AvatarUrl, &user.Email, &user.Password, &user.Status, &user.Blocked, &user.EmailConfirmed, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("no user found with this Id")
		}
		fmt.Println(err)
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (models.User, error) {
	row, err := r.DB.GetOneFrom(r.TableName, q.WhereOption{"username": username})
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	}
	var user models.User
	err = row.Scan(&user.UserId, &user.Firstname, &user.Lastname, &user.Username, &user.AvatarUrl, &user.Email, &user.Password, &user.Status, &user.Blocked, &user.EmailConfirmed, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("no user found with this username")
		}
		fmt.Println(err)
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	row, err := r.DB.GetOneFrom(r.TableName, q.WhereOption{"email": email})
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	}
	var user models.User
	err = row.Scan(&user.UserId, &user.Firstname, &user.Lastname, &user.Username, &user.AvatarUrl, &user.Email, &user.Password, &user.Status, &user.Blocked, &user.EmailConfirmed, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("no user found with this username")
		}
		fmt.Println(err)
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user models.User) error {
	err := r.DB.Update(r.TableName, user, q.WhereOption{"user_id": user.UserId})
	return err
}

func (r *UserRepository) DeleteUser(user models.User) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"user_id": user.UserId})
	return err
}

func (r *UserRepository) SaveUser(user models.User) error {
	err := r.DB.Insert(r.TableName, user)
	return err
}
