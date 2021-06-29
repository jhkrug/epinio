package testenv

import (
	"github.com/epinio/epinio/acceptance/helpers/proc"
	"github.com/epinio/epinio/helpers"

	. "github.com/onsi/gomega"
)

func SetupInClusterServices() {
	out, err := proc.Run(Root()+"/dist/epinio-linux-amd64 enable services-incluster", "", false)
	ExpectWithOffset(1, err).ToNot(HaveOccurred(), out)
	ExpectWithOffset(1, out).To(ContainSubstring("Beware, "))

	// Wait until classes appear
	EventuallyWithOffset(1, func() error {
		_, err = helpers.Kubectl("get clusterserviceclass mariadb")
		return err
	}, "5m").ShouldNot(HaveOccurred())

	// Wait until plans appear
	EventuallyWithOffset(1, func() error {
		_, err = helpers.Kubectl("get clusterserviceplan mariadb-10-3-22")
		return err
	}, "5m").ShouldNot(HaveOccurred())
}
