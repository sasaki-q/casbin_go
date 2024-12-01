package pkg

import (
	"os"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	jsonadapter "github.com/casbin/json-adapter/v2"
)

func NewCasbinEnforcer() (*casbin.Enforcer, error) {
	p, err := readFile("config/policy.json")
	if err != nil {
		return nil, err
	}

	m, err := readFile("config/model.conf")
	if err != nil {
		return nil, err
	}

	ms, err := model.NewModelFromString(string(m))
	if err != nil {
		return nil, err
	}

	return casbin.NewEnforcer(ms, jsonadapter.NewAdapter(&p))
}

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
