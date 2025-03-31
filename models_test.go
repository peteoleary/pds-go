package main

import (
	"strings"
	"testing"
)

func TestHashDid(t *testing.T) {
	did := "did:plc:example"
	expectedHash := "93aa4dfce0edcdbb4675901a956529df673e3c7f4686dbae6945582ffb5395a6"

	hash := hash_did(did)
	if hash != expectedHash {
		t.Errorf("Expected %s, got %s", expectedHash, hash)
	}
}

func TestGetActorDirectory(t *testing.T) {
	did := "did:plc:2yn32k65auyhjo2thnya3hlg"
	expectedDirectorySuffix := "35/did:plc:2yn32k65auyhjo2thnya3hlg"

	actor := Actor{
		did: did,
	}
	directory := actor.get_actor_directory()
	if !strings.HasSuffix(directory, expectedDirectorySuffix) {
		t.Errorf("Expected %s, got %s", expectedDirectorySuffix, directory)
	}
}
