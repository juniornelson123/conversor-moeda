package database

import (
	"testing"

	"github.com/juniornelson123/conversor-moeda/config/database"
)

func TestInitDB(t *testing.T) {
	resp, err := database.OpenDB("root", "root", "conversormoeda")
	if err != nil {
		t.Errorf(err.Error())
	}
	_, errBegin := resp.Begin()

	if errBegin != nil {

		t.Errorf(errBegin.Error())
	}

	defer database.CloseDB(resp)

}
