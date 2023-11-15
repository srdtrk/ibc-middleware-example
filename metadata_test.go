package example_test

import (
	"testing"

	ibcmock "github.com/cosmos/ibc-go/v8/testing/mock"
	"github.com/cosmosregistry/example"

	"github.com/stretchr/testify/require"
)

func TestMetadataFromVersion(t *testing.T) {
	var (
		version     string
		expMetadata example.Metadata
	)

	testCases := []struct {
		name     string
		malleate func()
		expErr   error
	}{
		{
			"success",
			func() {
				version = `{"example_version":"` + example.Version + `","app_version":"` + ibcmock.Version + `"}`
				expMetadata = example.NewMetadata(example.Version, ibcmock.Version)
			},
			nil,
		},
		{
			"failure: version does not contain example_version",
			func() {
				version = `{"invalid_version":"` + example.Version + `","app_version":"` + ibcmock.Version + `"}`
			},
			example.ErrInvalidChannelVersion,
		},
		{
			"failure: version does not contain app_version",
			func() {
				version = `{"example_version":"` + example.Version + `","invalid_version":"` + ibcmock.Version + `"}`
			},
			example.ErrInvalidChannelVersion,
		},
		{
			"failure: version is empty",
			func() {},
			example.ErrInvalidChannelVersion,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			version = ""
			expMetadata = example.Metadata{}

			tc.malleate()

			metadata, err := example.MetadataFromVersion(version)

			if tc.expErr == nil {
				require.NoError(t, err)
				require.Equal(t, expMetadata, metadata)
			} else {
				require.ErrorIs(t, err, tc.expErr)
				require.Empty(t, metadata)
			}
		})
	}
}
