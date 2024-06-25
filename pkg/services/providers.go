package services

import (
	"paragrafio/pkg/models"
)

type Providers struct {
	TransactionProvider models.TransactionProvider
	PageCache models.PageCache
	S3Storage models.S3Storage
}