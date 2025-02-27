package integration_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"

	"code.cloudfoundry.org/cf-networking-helpers/testsupport/metrics"
	"code.cloudfoundry.org/cf-networking-helpers/testsupport/ports"
	"code.cloudfoundry.org/policy-server/config"
	"code.cloudfoundry.org/policy-server/integration/helpers"
	testhelpers "code.cloudfoundry.org/test-helpers"
	. "github.com/onsi/ginkgo"
	ginkgoConfig "github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/onsi/gomega/types"
)

var (
	policyServerPath         string
	policyServerInternalPath string
	migrateDbPath            string
)

type policyServerPaths struct {
	Internal  string
	External  string
	MigrateDb string
}

var HaveName = func(name string) types.GomegaMatcher {
	return WithTransform(func(ev metrics.Event) string {
		return ev.Name
	}, Equal(name))
}

var HaveOriginAndName = func(origin, name string) types.GomegaMatcher {
	return SatisfyAll(
		WithTransform(func(ev metrics.Event) string {
			return ev.Name
		}, Equal(name)),
		WithTransform(func(ev metrics.Event) string {
			return ev.Origin
		}, Equal(origin)),
	)
}

var _ = helpers.MockCCServer
var _ = helpers.MockUAAServer

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var _ = SynchronizedBeforeSuite(func() []byte {
	var err error
	paths := policyServerPaths{}
	fmt.Fprint(GinkgoWriter, "building policy-server binary...")
	paths.External, err = gexec.Build("code.cloudfoundry.org/policy-server/cmd/policy-server", "-race")
	fmt.Fprint(GinkgoWriter, "done")
	Expect(err).NotTo(HaveOccurred())

	fmt.Fprint(GinkgoWriter, "building policy-server-internal binary...")
	paths.Internal, err = gexec.Build("code.cloudfoundry.org/policy-server/cmd/policy-server-internal", "-race")
	fmt.Fprint(GinkgoWriter, "done")
	Expect(err).NotTo(HaveOccurred())

	fmt.Fprint(GinkgoWriter, "building migrate-db binary...")
	paths.MigrateDb, err = gexec.Build("code.cloudfoundry.org/policy-server/cmd/migrate-db", "-race")
	fmt.Fprint(GinkgoWriter, "done")
	Expect(err).NotTo(HaveOccurred())

	data, err := json.Marshal(paths)
	Expect(err).NotTo(HaveOccurred())
	return data
}, func(data []byte) {
	var paths policyServerPaths
	err := json.Unmarshal(data, &paths)
	Expect(err).NotTo(HaveOccurred())

	policyServerPath = paths.External
	policyServerInternalPath = paths.Internal
	migrateDbPath = paths.MigrateDb

	rand.Seed(ginkgoConfig.GinkgoConfig.RandomSeed + int64(GinkgoParallelNode()))
})

var _ = SynchronizedAfterSuite(func() {}, func() {
	gexec.CleanupBuildArtifacts()
})

func configurePolicyServers(template config.Config, instances int) []config.Config {
	var configs []config.Config
	for i := 0; i < instances; i++ {
		conf := template
		conf.ListenPort = ports.PickAPort()
		conf.DebugServerPort = ports.PickAPort()
		configs = append(configs, conf)
	}
	return configs
}

func configureInternalPolicyServers(template config.InternalConfig, instances int) []config.InternalConfig {
	var configs []config.InternalConfig
	for i := 0; i < instances; i++ {
		conf := template
		conf.InternalListenPort = ports.PickAPort()
		conf.DebugServerPort = ports.PickAPort()
		conf.HealthCheckPort = ports.PickAPort()
		configs = append(configs, conf)
	}
	return configs
}

func startPolicyServers(configs []config.Config) []*gexec.Session {
	return startPolicyAndInternalServers(configs, nil)
}

func startPolicyAndInternalServers(configs []config.Config, internalConfigs []config.InternalConfig) []*gexec.Session {
	testhelpers.CreateDatabase(configs[0].Database)

	session := helpers.RunMigrationsPreStartBinary(migrateDbPath, configs[0])
	Eventually(session.Wait(TimeoutShort)).Should(gexec.Exit(0))

	var sessions []*gexec.Session
	for _, conf := range configs {
		sessions = append(sessions, helpers.StartPolicyServer(policyServerPath, conf))
	}

	for _, conf := range internalConfigs {
		sessions = append(sessions, helpers.StartInternalPolicyServer(policyServerInternalPath, conf))
	}
	return sessions
}

func stopPolicyServerExternalAndInternal(sessions []*gexec.Session, externalConfs []config.Config, internalConfs []config.InternalConfig) {
	for _, session := range sessions {
		session.Interrupt()
		Eventually(session, helpers.DEFAULT_TIMEOUT).Should(gexec.Exit())
	}
	testhelpers.RemoveDatabase(externalConfs[0].Database)
	testhelpers.RemoveDatabase(internalConfs[0].Database)
}

func stopPolicyServers(sessions []*gexec.Session, configs []config.Config) {
	for _, session := range sessions {
		session.Interrupt()
		Eventually(session, helpers.DEFAULT_TIMEOUT).Should(gexec.Exit())
	}
	testhelpers.RemoveDatabase(configs[0].Database)
}

func policyServerUrl(route string, confs []config.Config) string {
	conf := confs[rand.Intn(len(confs))]
	return fmt.Sprintf("http://%s:%d/networking/v1/%s", conf.ListenHost, conf.ListenPort, route)
}
