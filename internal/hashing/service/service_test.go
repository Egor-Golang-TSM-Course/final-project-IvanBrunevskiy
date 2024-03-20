package service

import (
	"testing"
)

func TestHashingService(t *testing.T) {
	service := NewHashingService()

	payload := "ivan"
	hash, _ := service.CreateHash(payload)

	expectedHash := "cd0b9452fc376fc4c35a60087b366f70d883fc901524daf1f122fbd319384f6a"
	if hash != expectedHash {
		t.Errorf("CreateHash() = %v, want %v", hash, expectedHash)
	}

	exists, _ := service.CheckHash(payload)
	if !exists {
		t.Errorf("CheckHash() = %v, want %v", exists, true)
	}

	retrievedHash, _ := service.GetHash(payload)
	if retrievedHash != hash {
		t.Errorf("GetHash() = %v, want %v", retrievedHash, hash)
	}

	_, err := service.CreateHash(payload)
	if err == nil {
		t.Error("Expected error when creating hash for existing payload, got nil")
	}
}
