package database

import "github.com/ResourceAPI/Core/database/filters"

type Database interface {
	Wipe()
	GetResources(name string, filters []filters.Filter) []map[string]interface{}
	StoreResources(name string, resources []map[string]interface{})
	DeleteResources(name string, filters []filters.Filter)
}
