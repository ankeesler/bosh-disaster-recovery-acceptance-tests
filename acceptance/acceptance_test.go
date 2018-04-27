package acceptance_test

import (
	"github.com/cloudfoundry-incubator/bosh-disaster-recovery-acceptance-tests/runner"
	"github.com/cloudfoundry-incubator/bosh-disaster-recovery-acceptance-tests/testcases"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("backing up bosh", func() {
	testCases := []runner.TestCase{
		testcases.ToyTestcase{},
	}

	runner.RunBoshDisasterRecoveryAcceptanceTests(config, testCases)
})
