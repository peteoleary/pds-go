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

func set_test_env(t *testing.T) {
	t.Setenv("PDS_DATA_DIRECTORY", "./testdata")
	t.Setenv("PDS_ACTOR_STORE_DIRECTORY", "./testdata/actors")
}

func TestGetRecordDatabaseLocation(t *testing.T) {
	set_test_env(t)

	did := "did:plc:2yn32k65auyhjo2thnya3hlg"
	expectedDirectorySuffix := "00/did:plc:"

	record := Record{
		actor_did: did,
	}
	directory := record.database_location()
	if !strings.HasSuffix(directory, expectedDirectorySuffix) {
		t.Errorf("Expected %s, got %s", expectedDirectorySuffix, directory)
	}
}
