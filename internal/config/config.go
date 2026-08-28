package config

import "os"

type Config struct {
	Address, Database, AccessSalt string
	MaxPhotos                     int
}

func Load() Config {
	c := Config{Address: ":8080", Database: "album.db", AccessSalt: "private", MaxPhotos: 1000}
	if v := os.Getenv("ALBUM_ADDR"); v != "" {
		c.Address = v
	}
	if v := os.Getenv("ALBUM_DB"); v != "" {
		c.Database = v
	}
	return c
}
func (c Config) Normalize() Config {
	if c.MaxPhotos < 1 {
		c.MaxPhotos = 1000
	}
	if c.Address == "" {
		c.Address = ":8080"
	}
	return c
}
func (c Config) Validate() error {
	if c.Database == "" {
		return os.ErrInvalid
	}
	return nil
}
