package hashid

import (
	"errors"
	"strings"

	"github.com/speps/go-hashids"
)

const MIN_HASH_ID_LENGTH = 6

var hid, _ = hashids.NewWithData(&hashids.HashIDData{
	Alphabet:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
	MinLength: MIN_HASH_ID_LENGTH,
})

func Encode(val int64) string {
	id, _ := hid.EncodeInt64(([]int64{val}))
	return id
}

func DecodeInt64(id string) (int64, error) {
	if id == "" {
		return 0, errors.New("HashId is empty")
	}

	res, err := hid.DecodeInt64WithError(id)
	if err != nil {
		return 0, err
	}
	if len(res) != 1 {
		return 0, errors.New("Hashid invalid")
	}
	return res[0], nil
}

func DecodeUrlParam(urlParam string) (int64, error) {
	if urlParam == "" {
		return 0, errors.New("UrlParam is empty")
	}

	x := strings.Split(urlParam, "-")
	if len(x) < 2 || len(x[len(x)-1]) < MIN_HASH_ID_LENGTH {
		return 0, errors.New("UrlParam is invalid")
	}

	return DecodeInt64(x[len(x)-1])
}
