package infrastructure

import "fmt"

type Environment string

const (
	EnvLocal      Environment = "local"
	EnvDevelop    Environment = "develop"
	EnvProduction Environment = "production"
	EnvTest       Environment = "test"
)

func ConvertEnvironment(rawEnvironment string) (*Environment, error) {
	environments := map[string]Environment{
		string(EnvTest):       EnvTest,
		string(EnvLocal):      EnvLocal,
		string(EnvDevelop):    EnvDevelop,
		string(EnvProduction): EnvProduction,
	}

	env, ok := environments[rawEnvironment]
	if !ok {
		return nil, fmt.Errorf("incorrect environment")
	}

	return &env, nil
}
