package handlers

import (
	"log"
)

// Products handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a new products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}
