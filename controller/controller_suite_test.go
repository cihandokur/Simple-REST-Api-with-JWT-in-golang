package controller_test

import (
	"testing"

	"github.com/cihandokur/simple_rest_api/config"
	"github.com/cihandokur/simple_rest_api/db"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestController(t *testing.T) {
	config.LoadConfiguration()
	db.New()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}
