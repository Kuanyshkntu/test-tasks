package service

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strconv"
	"test_tasks/models"
	"test_tasks/repository"
)

type service struct {
	repo                  repository.Repository
	ExternalApiHttpClient *http.Client
	logger                *zap.Logger
}

type Service interface {
	GetPeople(params map[string]string) ([]models.Person, error)
	AddPerson(person *models.Person) error
	UpdatePerson(person models.Person) error
	DeletePersonByID(id string) error
}

func NewService(
	repo repository.Repository,
) Service {
	return service{
		repo: repo,
	}
}

func (s service) AddPerson(person *models.Person) error {
	err := enrichPersonData(person)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = s.repo.AddPerson(person)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (s service) GetPeople(params map[string]string) ([]models.Person, error) {
	people, err := s.repo.GetPeople(params)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return people, nil
}

func (s service) UpdatePerson(person models.Person) error {
	err := s.repo.UpdatePerson(person)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (s service) DeletePersonByID(id string) error {
	err := s.repo.DeletePersonByID(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func enrichPersonData(person *models.Person) error {
	age, err := getEnrichedData("https://api.agify.io/?name="+person.Name, "age")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	iAge, err := strconv.Atoi(age)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	person.Age = iAge

	gender, err := getEnrichedData("https://api.genderize.io/?name="+person.Name, "gender")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	person.Gender = gender

	nationality, err := getEnrichedData("https://api.nationalize.io/?name="+person.Name, "nationality")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	person.Nationality = nationality
	return nil
}

func getEnrichedData(apiURL, key string) (res string, err error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	if key == "age" {
		age, ok := data[key].(float64)
		if !ok {
			return "", fmt.Errorf("Failed to get age from response")
		}
		res = fmt.Sprintf("%d", int(age))
	} else if key == "gender" {
		result, ok := data[key].(string)
		if !ok {
			return "", fmt.Errorf("Failed to get gender or nationality from response")
		}
		res = result
	} else {
		highestProbability := 0.0
		var selectedCountry string
		if countries, ok := data["country"].([]interface{}); ok {
			for _, country := range countries {
				if countryMap, ok := country.(map[string]interface{}); ok {
					if probability, ok := countryMap["probability"].(float64); ok {
						if probability > highestProbability {
							highestProbability = probability
							selectedCountry = countryMap["country_id"].(string)
						}
					}
				}
			}
		}
		res = selectedCountry
	}

	return res, nil
}
