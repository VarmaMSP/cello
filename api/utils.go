package api

import (
	"errors"

	"github.com/varmamsp/cello/model"
)

func GetId(i interface{}) (int64, error) {
	hashId, ok := i.(string)
	if !ok {
		return 0, errors.New("")
	}

	id, err := model.Int64FromHashId(hashId)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetIds(i interface{}) ([]int64, error) {
	x, ok := i.([]interface{})
	if !ok {
		return nil, errors.New("")
	}

	ids := make([]int64, len(x))
	for i := 0; i < len(x); i++ {
		hashId, ok := x[i].(string)
		if !ok {
			return nil, errors.New("")
		}

		id, err := model.Int64FromHashId(hashId)
		if err != nil {
			return nil, err
		}

		ids[i] = id
	}

	return ids, nil
}
