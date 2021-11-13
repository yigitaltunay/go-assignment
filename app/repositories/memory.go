package repositories

import (
	"fmt"
	"time"

	"github.com/yigitaltunay/go-assignment/app/models"
)

var MemoryDb = make(map[string]models.InMemoryModel)

type Repo struct{}

var R Repo

// In-Memory Insert Function
func (r Repo) MemoryDbInsert(data models.InMemoryModel) {
	MemoryDb[data.Key] = data
}

func (r Repo) MemoryDbFetch() map[string]models.InMemoryModel {
	return MemoryDb
}

// 10 minutes cache time
func (r Repo) MemoryDbFetchWithKey(key string) (models.InMemoryModel, error) {
	if val, ok := MemoryDb[key]; ok {
		if val.CreateDate.Add(time.Minute * 10).After(time.Now()) {
			return val, nil
		} else {
			delete(MemoryDb, key)
			return val, fmt.Errorf("expire date")
		}
	}
	return models.InMemoryModel{}, fmt.Errorf("key not found")
}

func (r Repo) IsExistsCurrentType(currentType string) error {
	switch currentType {
	case "EUR":
		return nil
	case "USD":
		return nil
	case "GBP":
		return nil
	default:
		return fmt.Errorf("currentType not found")
	}
}
