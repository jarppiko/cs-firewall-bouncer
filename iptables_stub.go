//go:build !linux
// +build !linux

package main

func newIPTables(config *bouncerConfig) (backend, error) {
	return nil, nil
}
