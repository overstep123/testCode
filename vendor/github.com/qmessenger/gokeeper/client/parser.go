package client

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/qmessenger/gokeeper/model"
)

func setStructField(itr interface{}, data map[string]model.ConfData) error {
	rfv := reflect.ValueOf(itr).Elem()
	for _, v := range data {
		field := rfv.FieldByName(v.StructKey)
		if !field.IsValid() {
			Stderr.Write([]byte(fmt.Sprintf("%s|gokeeper|setStructField|field invalid|%s \n", time.Now().String(), v.StructKey)))
			continue
		}
		typ := field.Type().String()
		if strings.Replace(typ, " ", "", -1) != v.Type {
			Stderr.Write([]byte(fmt.Sprintf("%s|gokeeper|setStructField|field type invalid|%s|%s|%s \n", time.Now().String(), v.StructKey, typ, v.Type)))
			continue
		}
		field.Set(reflect.ValueOf(v.Value))
	}
	return nil
}

func fill(rdata map[string]interface{}, sd model.StructData) (interface{}, error) {
	structInterface, ok := rdata[sd.Name]
	if !ok {
		return nil, errors.New("struct not load:" + sd.Name)
	}
	itr := reflect.New(reflect.TypeOf(structInterface)).Interface()
	if err := setStructField(itr, sd.Data); err != nil {
		return nil, err
	}

	return itr, nil
}
