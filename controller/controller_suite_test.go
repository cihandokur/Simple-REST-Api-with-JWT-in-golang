package controller_test

import (
	"testing"

	"github.com/cihandokur/devlab/config"
	"github.com/cihandokur/devlab/db"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestController(t *testing.T) {
	config.LoadConfiguration()
	db.New()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}
