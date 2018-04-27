package runner

import (
	"fmt"

	. "github.com/onsi/ginkgo"
)

func RunBoshDisasterRecoveryAcceptanceTests(config Config, testCases []TestCase) {
	It("backs up and restores bosh", func() {
		By("running the before backup step", func() {
			for _, testCase := range testCases {
				fmt.Println("Running the before backup step for " + testCase.Name())
				testCase.BeforeBackup()
			}
		})

		By("backing up", func() {
			fmt.Println("bbr director backup should run here")
			RunCommandSuccessfullyWithFailureMessage(
				"bbr director backup",
				fmt.Sprintf(
					"%s director --host %s --username %s --private-key-path %s backup --artifact-path %s",
					config.BBRBinaryPath,
					config.BOSH.Hostname,
					config.BOSH.SSHUsername,
					config.BOSH.SSHPrivateKeyPath,
					config.ArtifactPath,
				),
			)
		})

		By("running the after backup step", func() {
			for _, testCase := range testCases {
				fmt.Println("Running the after backup step for " + testCase.Name())
				testCase.AfterBackup()
			}
		})

		By("restoring", func() {
			fmt.Println("bbr director restore should run here")
		})

		By("running the after restore step", func() {
			for _, testCase := range testCases {
				fmt.Println("Running the after restore step for " + testCase.Name())
				testCase.AfterRestore()
			}
		})

		By("cleaning up", func() {
			for _, testCase := range testCases {
				fmt.Println("Running the cleanup step for " + testCase.Name())
				testCase.Cleanup()
			}
		})
	})
}
