package database

import (
	"fmt"
	"os"
)

type EnvBase struct {
}

func NewEnvBase() Database {
	return &EnvBase{}
}

func (d *EnvBase) CountAll() (int, error) {
	return 1, nil
}

func (d *EnvBase) DeviceTokenByKey(key string) (string, error) {
	if key == os.Getenv("BARK_KEY") {
		return os.Getenv("BARK_DEVICE_TOKEN"), nil
	}
	return "nil", fmt.Errorf("key not found")
}

func (d *EnvBase) SaveDeviceTokenByKey(key, token string) (string, error) {
	return os.Getenv("BARK_KEY"), nil
}

func (d *EnvBase) Close() error {
	return nil
}
