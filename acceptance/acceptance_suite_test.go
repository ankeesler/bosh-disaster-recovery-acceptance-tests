package acceptance_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"io/ioutil"
	"testing"

	"github.com/cloudfoundry-incubator/bosh-disaster-recovery-acceptance-tests/runner"
)

var config runner.Config

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acceptance Suite")

	BeforeSuite(func() {
		artifactDirPath, err := ioutil.TempDir("", "b-drats")
		Expect(err).NotTo(HaveOccurred())

		config = runner.Config{
			BOSH: runner.BOSHConfig{
				Hostname:          mustHaveEnv("BOSH_HOSTNAME"),
				SSHUsername:       mustHaveEnv("BOSH_SSH_USERNAME"),
				SSHPrivateKeyPath: mustHaveEnv("BOSH_SSH_PRIVATE_KEY_PATH"),
			},
			BBRBinaryPath: mustHaveEnv("BBR_BINARY_PATH"),
			ArtifactPath:  artifactDirPath,
		}
	})
}

func mustHaveEnv(keyname string) string {
	val := os.Getenv(keyname)
	if val == "" {
		panic(fmt.Sprintf("Env var %s not set", keyname))
	}
	return val
}
