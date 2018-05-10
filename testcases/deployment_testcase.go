package testcases

import (
	"fmt"
	"github.com/cloudfoundry-incubator/bosh-disaster-recovery-acceptance-tests/fixtures"
	. "github.com/cloudfoundry-incubator/bosh-disaster-recovery-acceptance-tests/runner"
)

type DeploymentTestcase struct{}

func (t DeploymentTestcase) Name() string {
	return "deployment-testcase"
}

func (t DeploymentTestcase) BeforeBackup(config Config) {
	RunCommandSuccessfullyWithFailureMessage("deploy syslog",
		fmt.Sprintf("bosh "+
			"--environment=%s "+
			"--client=%s "+
			"--client-secret=%s "+
			"--ca-cert=%s "+
			"--deployment=syslog "+
			"deploy %s -n",
			config.BOSH.Host,
			config.BOSH.Client,
			config.BOSH.ClientSecret,
			config.BOSH.CACertPath,
			fixtures.Path("syslog-manifest.yml"),
		))
}

func (t DeploymentTestcase) AfterBackup(config Config) {}

func (t DeploymentTestcase) AfterRestore(config Config) {}

func (t DeploymentTestcase) Cleanup(config Config) {}
