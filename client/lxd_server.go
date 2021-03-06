package lxd

import (
	"github.com/lxc/lxd/shared/api"
)

// Server handling functions

// GetServer returns the server status as a Server struct
func (r *ProtocolLXD) GetServer() (*api.Server, string, error) {
	server := api.Server{}

	// Return the cached entry if present
	if r.server != nil {
		return r.server, "", nil
	}

	// Fetch the raw value
	etag, err := r.queryStruct("GET", "", nil, "", &server)
	if err != nil {
		return nil, "", err
	}

	// Add the value to the cache
	r.server = &server

	return &server, etag, nil
}

// UpdateServer updates the server status to match the provided Server struct
func (r *ProtocolLXD) UpdateServer(server api.ServerPut, ETag string) error {
	// Send the request
	_, _, err := r.query("PUT", "", server, ETag)
	if err != nil {
		return err
	}

	return nil
}

// HasExtension returns true if the server supports a given API extension
func (r *ProtocolLXD) HasExtension(extension string) bool {
	for _, entry := range r.server.APIExtensions {
		if entry == extension {
			return true
		}
	}

	return false
}
