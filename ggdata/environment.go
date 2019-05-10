package ggdata

import (
	"github.com/vacovsky/gogrowggmodels"
)

// GetLatestEnvironmentCheck : return to the called the lastest environment DB entry
func GetLatestEnvironmentCheck() ggmodels.Environment {

	result := ggmodels.Environment{}
	Service().Last(&result)

	return result
}

// SaveEnvironment : save the data of the most recent environment check to the DB
func SaveEnvironment(env ggmodels.Environment) error {
	Service().Save(&env)
	return nil
}
