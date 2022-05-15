package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SQLiteWorker struct {
	DB *gorm.DB
}

//Создание новой записи User в базе данных с использованием хэширования пароля
func (worker *SQLiteWorker) CreateUser(user User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Passhash), 14)
	if err != nil {
		return err
	}
	user.Passhash = string(bytes)

	result := worker.DB.Create(&user)

	return result.Error
}

//Получение записи пользователя из базы данных по значению Mail
func (worker *SQLiteWorker) GetUser(mail string) (User, error) {
	var user User
	result := worker.DB.Where("Mail = ?", mail).First(&user)
	return user, result.Error
}

//Обновление записи в базе данных
func (worker *SQLiteWorker) UpdateUser(user User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Passhash), 14)
	if err != nil {
		return err
	}
	user.Passhash = string(bytes)

	result := worker.DB.Model(&User{}).Where("Mail = ?", user.Mail).Updates(user)

	return result.Error
}

//Удаление записи базы данных записи по значению Mail
func (worker *SQLiteWorker) DeleteUser(mail string) error {
	result := worker.DB.Where("Mail = ?", mail).Delete(&User{})
	return result.Error
}

//Проверка наличия записи в базе данных 
//и соответствия переданного хэша хешу записанному в базе данных
func (worker *SQLiteWorker) CheckUser(user User) error {
	var checkUser User
	worker.DB.Where("Mail = ?", user.Mail).First(&checkUser)

	err := bcrypt.CompareHashAndPassword([]byte(checkUser.Passhash), []byte(user.Passhash))
	
	return err
}