package functional

import (
	"os"
	"testing"

	"k8s.io/client-go/rest"

	"github.com/integr8ly/integreatly-operator/test/common"
	runtimeConfig "sigs.k8s.io/controller-runtime/pkg/client/config"
)

func TestIntegreatly(t *testing.T) {
	config, err := runtimeConfig.GetConfig()
	config.Impersonate = rest.ImpersonationConfig{
		UserName: "system:admin",
		Groups:   []string{"system:authenticated"},
	}
	if err != nil {
		t.Fatal(err)
	}

	t.Run("API Managed Multi-AZ Tests", func(t *testing.T) {
		// Do not execute these tests unless MULTIAZ is set to true
		if os.Getenv("MULTIAZ") != "true" {
			t.Skip("Skipping Multi-AZ tests as MULTIAZ env var is not set to true")
		}

		common.RunTestCases(MULTIAZ_TESTS, t, config)
	})

	t.Run("Integreatly Destructive Tests", func(t *testing.T) {
		// Do not execute these tests unless DESTRUCTIVE is set to true
		if os.Getenv("DESTRUCTIVE") != "true" {
			t.Skip("Skipping Destructive tests as DESTRUCTIVE env var is not set to true")
		}

		common.RunTestCases(common.DESTRUCTIVE_TESTS, t, config)
	})
}
