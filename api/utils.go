package api

import (
	"github.com/varmamsp/cello/model"
)

func GetId(i interface{}) (int64, *model.AppError) {
	hashId, ok := i.(string)
	if !ok {
		return 0, model.NewAppError("api.get_id", "invalid_body_param", 400, nil)
	}

	id, err := model.Int64FromHashId(hashId)
	if err != nil {
		return 0, model.NewAppError("api.get_id", "invalid_body_param", 400, nil)
	}

	return id, nil
}

func GetIds(i interface{}) ([]int64, *model.AppError) {
	x, ok := i.([]interface{})
	if !ok {
		return []int64{}, nil
	}

	ids := make([]int64, len(x))
	for i := 0; i < len(x); i++ {
		tmp, err := GetId(x[i])
		if err != nil {
			return nil, err
		}

		ids[i] = tmp
	}

	return ids, nil
}
