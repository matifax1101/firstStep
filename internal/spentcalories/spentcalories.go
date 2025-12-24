package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	values := strings.Split(data, ",")
	if len(values) != 3 {
		return 0, "", 0, fmt.Errorf("неверный формат тренировки")
	}

	// шаги
	steps, err := strconv.Atoi(values[0])
	if err != nil {
		return 0, "", 0, err
	}

	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	// продолжительность
	duration, err := time.ParseDuration(values[2])
	if err != nil {
		return 0, "", 0, err
	}

	return steps, values[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0
	}

	stepLength := height * stepLengthCoefficient
	return (float64(steps) * stepLength) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	dist := distance(steps, height)
	return dist / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)

	var calories float64

	switch activity {
	case "Ходьба":
		calories, err = WalkingSpentCalories(
			steps,
			weight,
			height,
			duration,
		)
	case "Бег":
		calories, err = RunningSpentCalories(
			steps,
			weight,
			height,
			duration,
		)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity,
		duration.Hours(),
		dist,
		speed,
		calories,
	), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("некорректные входные данные")
	}

	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()

	return (weight * speed * minutes) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	calories, err := RunningSpentCalories(
		steps,
		weight,
		height,
		duration,
	)
	if err != nil {
		return 0, err
	}

	return calories * walkingCaloriesCoefficient, nil
}
