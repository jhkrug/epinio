package acceptance_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/epinio/epinio/acceptance/testenv"
	"github.com/onsi/ginkgo/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(FailWithReport)
	RunSpecs(t, "Acceptance Suite")
}

var (
	nodeSuffix, nodeTmpDir string
	env                    testenv.EpinioEnv
	epinioBinary           string
)

var _ = SynchronizedBeforeSuite(func() []byte {
	fmt.Printf("I'm running on runner = %s\n", os.Getenv("HOSTNAME"))

	testenv.SetRoot("../..")
	testenv.SetupEnv()

	if err := testenv.CheckDependencies(); err != nil {
		panic("Missing dependencies: " + err.Error())
	}

	fmt.Printf("Compiling Epinio on node %d\n", config.GinkgoConfig.ParallelNode)
	testenv.BuildEpinio()
	testenv.CreateRegistrySecret()

	epinioBinary = fmt.Sprintf("/dist/epinio-%s-%s", runtime.GOOS, runtime.GOARCH)

	return []byte(strconv.Itoa(int(time.Now().Unix())))
}, func(randomSuffix []byte) {
	var err error
	testenv.SetRoot("../..")

	nodeSuffix = fmt.Sprintf("%d-%s",
		config.GinkgoConfig.ParallelNode, string(randomSuffix))
	nodeTmpDir, err = ioutil.TempDir("", "epinio-"+nodeSuffix)
	if err != nil {
		panic("Could not create temp dir: " + err.Error())
	}

	Expect(os.Getenv("KUBECONFIG")).ToNot(BeEmpty(), "KUBECONFIG environment variable should not be empty")

	env = testenv.New(nodeTmpDir, testenv.Root())
})

var _ = SynchronizedAfterSuite(func() {
	if !testenv.SkipCleanup() {
		fmt.Printf("Deleting tmpdir on node %d\n", config.GinkgoConfig.ParallelNode)
		testenv.DeleteTmpDir(nodeTmpDir)
	}
}, func() { // Runs only on one node after all are done
	if testenv.SkipCleanup() {
		fmt.Printf("Found '%s', skipping all cleanup", testenv.SkipCleanupPath())
	} else {
		// Delete left-overs no matter what
		defer func() { _, _ = testenv.CleanupTmp() }()
	}
})

var _ = AfterEach(func() {
	testenv.AfterEachSleep()
})

func FailWithReport(message string, callerSkip ...int) {
	// NOTE: Use something like the following if you need to debug failed tests
	// fmt.Println("\nA test failed. You may find the following information useful for debugging:")
	// fmt.Println("The cluster pods: ")
	// out, err := helpers.Kubectl("get pods --all-namespaces")
	// if err != nil {
	// 	fmt.Print(err.Error())
	// } else {
	// 	fmt.Print(out)
	// }

	// Ensures the correct line numbers are reported
	Fail(message, callerSkip[0]+1)
}