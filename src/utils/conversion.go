package utils

import "strconv"

func AtoI(saisie string) int {
	value, err := strconv.Atoi(saisie)
	if err != nil {
		value = 0
	}
	return value
}

func AtoI64(saisie string) int64 {
	value, err := strconv.ParseInt(saisie, 10, 64)
	if err != nil {
		value = 0
	}
	return value
}

func ItoA(value int) string {
	return strconv.Itoa(value)
}

func I64toA(value int64) string {
	return strconv.FormatInt(int64(value), 10)
}

func AtoF(saisie string) float32 {
	value, err := strconv.ParseFloat(saisie, 32)
	if err != nil {
		value = 0
	}
	return float32(value)
}

func FtoA(value float32) string {
	if value < 0.1 {
		return strconv.FormatFloat(float64(value), 'f', 2, 32)
	} else {
		return strconv.FormatFloat(float64(value), 'f', 1, 32)
	}

}

func AtoB(saisie string) bool {
	if saisie == "1" {
		return true
	} else {
		return false
	}

}
