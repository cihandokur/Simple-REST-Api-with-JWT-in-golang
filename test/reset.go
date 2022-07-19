package test

import (
	"fmt"

	"github.com/cihandokur/devlab/db"
	. "github.com/onsi/gomega"
)

func ResetAll() {
	ResetUserTable()
}

func ResetUserTable() {

	fmt.Println("called ResetUserTable")

	cmd := "TRUNCATE users;"
	res := db.DB.Exec(cmd)
	Expect(res.Error).NotTo(HaveOccurred())
}
