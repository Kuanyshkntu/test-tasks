package repository

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"strconv"
	"test_tasks/models"
)

type repo struct {
	Db *gorm.DB
}

type Repository interface {
	GetPeople(params map[string]string) ([]models.Person, error)
	AddPerson(person *models.Person) error
	UpdatePerson(person models.Person) error
	DeletePersonByID(id string) error
}

func NewRepository(db *gorm.DB) Repository {
	return repo{Db: db}
}

func (f repo) AddPerson(person *models.Person) error {
	return f.Db.Create(&person).Error
}

func (f repo) GetPeople(params map[string]string) ([]models.Person, error) {
	query := `	
		select id, name, surname, coalesce(patronymic, ''), age, gender, nationality
	from people
		where 1=1
`
	if params["name"] != "" {
		query += fmt.Sprintf(" AND name = %s", params["name"])
	}
	if params["surname"] != "" {
		query += fmt.Sprintf(" AND name = %s", params["surname"])
	}
	if params["nationality"] != "" {
		query += fmt.Sprintf(" AND name = %s", params["nationality"])
	}

	limit, ok := params["limit"]
	if ok {
		if page, ok := params["page"]; ok {
			query += " LIMIT " + limit
			pageInt, err := strconv.Atoi(page)
			if err != nil {
				return nil, err
			}
			limitInt, err := strconv.Atoi(limit)
			if err != nil {
				return nil, err
			}
			offset := (pageInt - 1) * limitInt
			query += fmt.Sprintf(" OFFSET %d", offset)
		}
	}

	rows, err := f.Db.Raw(query).Rows()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var people []models.Person
	for rows.Next() {
		var person models.Person
		err = rows.Scan(
			&person.ID,
			&person.Name,
			&person.Surname,
			&person.Patronymic,
			&person.Age,
			&person.Gender,
			&person.Nationality,
		)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		people = append(people, person)
	}
	return people, nil
}

//func (f repo) GetByID(id string) (person models.Person, err error) {
//	query := `
//			select id, name, surname, coalesce(patronymic, ''), age, gender, nationality
//			from people
//			where id = $1
// 			`
//
//	err = f.Db.Select(query, id).Row().Scan(&person.ID, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.Nationality)
//	if err != nil {
//		log.Println(err.Error())
//		return person, err
//	}
//	return person, err
//}

func (f repo) UpdatePerson(person models.Person) error {
	query := `
		update people
		set 
		    name = COALESCE($2,name),
		    surname = COALESCE($3,surname),
		    patronymic = COALESCE($4,patronymic)
		where id = $1
`
	err := f.Db.Exec(query, person.ID, person.Name, person.Surname, person.Patronymic).Error
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (f repo) DeletePersonByID(id string) error {
	query := `
		delete
		from people
		where id = $1
`
	err := f.Db.Exec(query, id).Error
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
