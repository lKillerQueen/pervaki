package service

import (
	"fmt"
	"pervaki/model"
)

type AnimalService struct{}

func NewAnimalService() AnimalService {
	return AnimalService{}
}

func (s AnimalService) SoundByZoo(animal model.Zoo) (string, error) {
	switch animal {
	case model.Cat:
		return "Мяу-мяу", nil
	case model.Dog:
		return "Гав-гав", nil
	case model.Woman:
		return "Прости, ты мне просто друг", nil
	}

	return "", fmt.Errorf("нет такого животного: %s", animal)
}
