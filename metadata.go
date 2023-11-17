package example

import (
	"encoding/json"
	"strings"

	errorsmod "cosmossdk.io/errors"
)

// NewMetadata returns a new Metadata instance
func NewMetadata(exampleVersion, appVersion string) Metadata {
	return Metadata{
		ExampleVersion: exampleVersion,
		AppVersion:     appVersion,
	}
}

// MetadataFromVersion attempts to parse the given string into a fee version Metadata,
// an error is returned if it fails to do so.
func MetadataFromVersion(version string) (Metadata, error) {
	var metadata Metadata
	err := json.Unmarshal([]byte(version), &metadata)
	if err != nil {
		return Metadata{}, errorsmod.Wrapf(ErrInvalidChannelVersion, "failed to unmarshal metadata from version: %s", version)
	}

	err = validateMetadata(metadata)
	if err != nil {
		return Metadata{}, err
	}

	return metadata, nil
}

// validateMetadata validates the given metadata
func validateMetadata(metadata Metadata) error {
	if strings.TrimSpace(metadata.AppVersion) == "" {
		return errorsmod.Wrapf(ErrInvalidChannelVersion, "metadata app_version cannot be empty")
	}

	if strings.TrimSpace(metadata.ExampleVersion) == "" {
		return errorsmod.Wrapf(ErrInvalidChannelVersion, "metadata example_version cannot be empty")
	}

	return nil
}
