package postgres

import (
	"fmt"
)

func (service *DbService) execute(query string, args ...interface{}) error {
	exec, err := service.connection.Exec(query, args)
	if err != nil {
		return err
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		return fmt.Errorf("error executing query='%s' with args='%s'", query, args)
	}

	return nil
}
