package utils

import (
	"errors"
	"fmt"
)

type Cache struct {
	m map[string]string
}

func Construct() *Cache {
	return &Cache{m: make(map[string]string)} // Правильное определение мапы
}
func Set(mape *Cache, key, value string) *Cache {
	mape.m[key] = value
	return mape
}
func Exist(mape *Cache, key string) error {
	var err error
	if mape.m[key] != "" {
		err = errors.New("Key already exist")
		return err
	}
	return err
}
func Get(mape *Cache, key string) string {
	if Exist(mape, key) != nil {
		return mape.m[key]
	}
	return "Key doesn't exist"
}
func Del(mape *Cache, key string) {
	delete(mape.m, key)
}
func Getbykeys(mape *Cache, key []string) []string {
	var temp []string
	for k, v := range mape.m {
		for i := 0; i < len(key); i++ {
			if key[i] == k {
				temp = append(temp, v)
				//fmt.Printf("key[%s] value[%s]\n", k, v)
			}
		}
	}
	return temp
}

func main() {
	mape := Construct()
	Set(mape, "1", "One")
	Set(mape, "2", "Two")
	Set(mape, "3", "Three")
	Exist(mape, "1")
	Exist(mape, "4")
	var keys []string = []string{"1", "2"}
	fmt.Println(Get(mape, "3"))
	fmt.Println(Get(mape, "4"))
	fmt.Printf("%+v", Getbykeys(mape, keys))
	Del(mape, "1")
	fmt.Println(mape.m)
}
