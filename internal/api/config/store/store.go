package store 

import (
	"context"
	"errors"
	"time"
)

var (
	ErrNotFound = errors.New("link not found")
	ErrCodeTaken = errors.New("code already taken")
)

type Link struct {
	Code	  string
	URL 	  string 
	CreatedAT time.Time
}

type Store interface {

	Create(ctx context.Context, code, url string) error 

	Get(ctx context.Context, code string) (Link, error)

	Ping(ctx context.Context) error 
}