package errors

import (
	"errors"
)

var (
	ErrObjectNotFound     = errors.New("object not found inside the database")
	ErrWrongDBEdition     = errors.New("the Portainer database is set for Portainer Business Edition, please follow the instructions in our documentation to downgrade it: https://documentation.portainer.io/v2.0-be/downgrade/be-to-ce/")
	ErrDBImportFailed     = errors.New("importing backup failed")
	ErrDatabaseIsUpdating = errors.New("database is currently in updating state. Failed prior upgrade. Please restore from backup or delete the database and restart Portainer")
)
