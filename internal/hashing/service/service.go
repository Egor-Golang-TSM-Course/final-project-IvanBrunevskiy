package service

import (
	"crypto/sha256"
	"fmt"
)

type HashingService struct {
	store map[string]string
}

func NewHashingService() *HashingService {
	return &HashingService{
		store: make(map[string]string),
	}
}

func (s *HashingService) CheckHash(payload string) (bool, error) {
	_, exists := s.store[payload]
	return exists, nil
}

func (s *HashingService) GetHash(payload string) (string, error) {
	hash, exists := s.store[payload]
	if !exists {
		return "", fmt.Errorf("hash not found for payload: %s", payload)
	}
	return hash, nil
}

func (s *HashingService) CreateHash(payload string) (string, error) {
	if _, exists := s.store[payload]; exists {
		return "", fmt.Errorf("hash already exists for payload: %s", payload)
	}

	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	s.store[payload] = hash

	return hash, nil
}
