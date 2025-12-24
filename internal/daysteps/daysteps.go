package daysteps

import (
	"time"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	values := strings.Split(data, ",")
	if len(values) != 2 {
		return 0, 0, fmt.Errorf("неверный формат данных")
}
// парсим шаги
	steps, err := strconv.Atoi(values[0])
	if err != nil {
		return 0, 0, err
	}
// колличество шагов - проверка
	if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть больше 0")
	}
// продолжительность
	duration, err := time.ParseDuration(values[1])
	if err != nil {
		return 0, 0, err
	}
	if duration <= 0 {
		return 0, 0, fmt.Errorf("продолжительность должна быть больше 0")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	// 2. дистанция 
	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	// 3. расчёт калории
	calories, err := spentcalories.WalkingSpentCalories(
		steps,
		weight,
		height,
		duration,
	)
	if err != nil {
		log.Println(err)
		return ""
	}

	// 4. вывод результата
	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps,
		distanceKm,
		calories,
	)
}

