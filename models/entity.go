package models

type Person struct {
	ID          int     `json:"id" gorm:"id"`
	Name        string  `json:"name" gorm:"name"`
	Surname     string  `json:"surname" gorm:"surname"`
	Patronymic  *string `json:"patronymic" gorm:"patronymic"`
	Age         int     `json:"age" gorm:"age"`
	Gender      string  `json:"gender" gorm:"gender"`
	Nationality string  `json:"nationality" gorm:"nationality"`
}

type MessageType struct {
	//Message - локализованный текст ошибки
	Message string `json:"message"`
	//Description - err.Error()
	Description string `json:"description"`
}
