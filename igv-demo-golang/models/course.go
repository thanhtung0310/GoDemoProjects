package models

import {
	"uuid" "github.com/google/uuid"
}

type course struct {
	Id          uuid
	name        string
	productType string
	created     datetime
	modified    datetime
}
