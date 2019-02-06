package testcases

import (
	"fmt"
	. "github.com/cloudfoundry-incubator/bosh-disaster-recovery-acceptance-tests/runner"
	. "github.com/onsi/gomega"
	"time"
)

type TruncateDBBlobstoreTestcase struct{}

func (t TruncateDBBlobstoreTestcase) Name() string {
	return "truncate_db_blobstore_testcase"
}

func (t TruncateDBBlobstoreTestcase) BeforeBackup(config Config) {
	monitStop(config, "director")
	monitStop(config, "uaa")
	monitStop(config, "credhub")
	monitStop(config, "blobstore_nginx")
	monitStop(config, "postgres")

	RunCommandInDirectorVMSuccessfullyWithFailureMessage(
		"truncate db/blobstore",
		config,
		"sudo rm -rf /var/vcap/store/{blobstore,director,postgres*}",
	)

	RunCommandInDirectorVMSuccessfullyWithFailureMessage(
		"pre-start all jobs",
		config,
		"sudo bash -c",
		"'for pre in $(ls /var/vcap/jobs/**/bin/pre-start); do $pre; done'",
	)

	fmt.Println("Sleeping for 20 seconds")
	time.Sleep(time.Second * 20)

	monitStart(config, "postgres")
	monitStart(config, "blobstore_nginx")
	monitStart(config, "uaa")
	monitStart(config, "credhub")
	monitStart(config, "director")

	Eventually(func() int {
		session := RunBoshCommand(
			"bosh env",
			config,
			"env",
		)
		return session.Wait().ExitCode()
	}).Should(Equal(0))

	Expect("if we got this far, we've passed").To(Equal(""))
}

func (t TruncateDBBlobstoreTestcase) AfterBackup(config Config) {

}

func (t TruncateDBBlobstoreTestcase) AfterRestore(config Config) {

}

func (t TruncateDBBlobstoreTestcase) Cleanup(config Config) {

}
