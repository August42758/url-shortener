package repository

import "errors"

var ErrUrlAlreadyExists error = errors.New("Такая ссылка уже существует")
var ErrUrlNotFound error = errors.New("Такой ссылки не существует")
