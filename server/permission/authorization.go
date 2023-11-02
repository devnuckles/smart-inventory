package permission

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/Tonmoy404/Smart-Inventory/logger"
	"github.com/Tonmoy404/Smart-Inventory/service"
)

type Authorization interface {
	service.Authorization
}

type authorization struct {
	Roles map[string][]permission
}

func NewAuthorization() Authorization {
	var data authorizationData
	if err := readJSONFile("permission/data.json", &data); err != nil {
		log.Panic("cannot read json file: ", err)
	}

	log.Printf("Authorization: %+v\n", data)

	return &authorization{
		Roles: map[string][]permission{
			"super-admin": data.Roles.SuperAdmin,
			"admin":       data.Roles.Admin,
			"reporter":    data.Roles.Reporter,
			"user":        data.Roles.User,
		},
	}
}

func (a *authorization) IsPermitted(ctx context.Context, role, action, object string) bool {
	if permissions, ok := a.Roles[role]; ok {
		for _, permission := range permissions {
			if permission.Action == action && permission.Object == object {
				return true // permission granted
			}
		}
	}

	logger.Info(ctx, "Role does not exist", role)

	return false
}

func readJSONFile(filename string, result interface{}) error {
	// open the JSON file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// read the JSON content
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// unmarshal JSON into the provided struct
	if err := json.Unmarshal([]byte(byteValue), result); err != nil {
		return err
	}

	return nil
}
