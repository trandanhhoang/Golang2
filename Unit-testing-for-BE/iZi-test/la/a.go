package la

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func EncodeFillInAnswer(answer string) string {
	char := regexp.MustCompile(`[^[[:blank:]]]*?`)
	return char.ReplaceAllString(answer, "*")
}

func MakeRange(start, end int) []int {
	cap := int(math.Abs(float64(end - start)))
	if cap == 0 {
		return []int{}
	}

	slice := make([]int, 0, cap)
	step := (end - start) / cap
	for i := start; i != end; i += step {
		slice = append(slice, i)
	}

	return slice
}

func ShuffleIndexes(sample int, population int) ([]int, error) {
	if sample > population {
		return nil, errors.New("unable to shuffle indexes")
	}

	generated := map[int]bool{}
	shuffledIndexes := make([]int, 0, sample)
	rand.Seed(time.Now().UnixNano())
	for {
		newIndex := rand.Int() % population
		if _, ok := generated[newIndex]; ok {
			continue
		}

		generated[newIndex] = true
		shuffledIndexes = append(shuffledIndexes, newIndex)
		if len(shuffledIndexes) == cap(shuffledIndexes) {
			break
		}
	}

	return shuffledIndexes, nil
}

var ZeroBool bool
var ZeroInt int
var ZeroString string
var ZeroMap map[string]interface{}

func GetBool(value interface{}) (bool, error) {
	switch t := value.(type) {
	case int:
		fmt.Printf("HOANG TEST FOR SURE")
		return true, nil
	case []byte:
		b, err := strconv.ParseBool(string(value.([]byte)))
		if err != nil {
			return ZeroBool, err
		}
		return b, nil
	case bool:
		return value.(bool), nil
	case string:
		b, err := strconv.ParseBool(value.(string))
		if err != nil {

			fmt.Println("True")
			fmt.Println(ZeroBool)
			fmt.Println(b)
			return ZeroBool, err
		}
		return b, nil
	default:
		return ZeroBool, fmt.Errorf("cannot convert %v of type %v to bool", value, t)
	}
}

func GetString(value interface{}) (string, error) {
	switch t := value.(type) {
	case string:
		return value.(string), nil
	default:
		b, err := json.Marshal(value)
		if err != nil {
			return ZeroString, fmt.Errorf("error marshaling %v of type %v: %v", value, t, err)
		}
		return string(b), nil
	}
}
