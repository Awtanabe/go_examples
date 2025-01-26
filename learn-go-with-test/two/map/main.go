package map_test

import "errors"


type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", errors.New("not found")
	}
	return val, nil
}

func (d Dictionary) Add(key, val string) error {
	if key == "" || val == "" {
		return errors.New("不足")
	}
	d[key] = val
  return nil
}