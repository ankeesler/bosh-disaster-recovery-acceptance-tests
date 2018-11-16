package testcases

import (
	. "github.com/cloudfoundry-incubator/bosh-disaster-recovery-acceptance-tests/runner"
)

type DeploymentTestcase struct{}

func (t DeploymentTestcase) Name() string {
	return "deployment_testcase"
}

func (t DeploymentTestcase) BeforeBackup(config Config) {
}

func (t DeploymentTestcase) AfterBackup(config Config) {
}

func (t DeploymentTestcase) AfterRestore(config Config) {
}

func (t DeploymentTestcase) Cleanup(config Config) {
}
