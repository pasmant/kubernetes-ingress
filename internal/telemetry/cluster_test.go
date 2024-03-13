package telemetry_test

import (
	"context"
	"testing"

	"github.com/nginxinc/kubernetes-ingress/internal/telemetry"
	apiCoreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestNodeCountInAClusterWithThreeNodes(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, node1, node2, node3)

	got, err := c.NodeCount(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestNodeCountInAClusterWithOneNode(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, node1)
	got, err := c.NodeCount(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	want := 1
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestClusterIDRetrievesK8sClusterUID(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, node1, kubeNS)

	got, err := c.ClusterID(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "329766ff-5d78-4c9e-8736-7faad1f2e937"
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestClusterIDErrorsOnNotExistingService(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, node1)
	_, err := c.ClusterID(context.Background())
	if err == nil {
		t.Error("want error, got nil")
	}
}

func TestK8sVersionRetrievesClusterVersion(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, node1)
	got, err := c.ClusterVersion()
	if err != nil {
		t.Fatal(err)
	}

	want := "v1.29.2"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestAWSPlatformDeterminesOwnName(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, nodeAWS)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "aws"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestAzurePlatformDeterminesOwnName(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, nodeAzure)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "azure"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestGCPPlatformDeterminesOwnName(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, nodeGCP)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "gce"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestKindPlatformDeterminesOwnName(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, nodeKind)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "kind"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestVSpherePlatformDeterminesOwnName(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, nodeVSphere)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "vsphere"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestK3SPlatformDeterminesOwnName(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, nodeK3S)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "k3s"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestIBMCloudPlatformDeterminesOwnName(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, nodeIBMCloud)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "ibmcloud"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestIBMPowerPlatformDeterminesOwnName(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, nodeIBMPowerVS)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "ibmpowervs"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestCloudStackPlatformDeterminesOwnName(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, nodeCloudStack)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "cloudstack"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestOpenStackPlatformDeterminesOwnName(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, nodeOpenStack)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "openstack"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestDigitalOceanPlatformDeterminesOwnName(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, nodeDigitalOcean)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "digitalocean"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestEquinixMetallPlatformDeterminesOwnName(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, nodeEquinixMetal)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "equinixmetal"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestAlibabaPlatformDeterminesOwnName(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, nodeAlibaba)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "alicloud"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestPlatformLookupOnMissingPlatformIDField(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, node1)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "other"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestPlatformLookupOnMalformedPlatformIDField(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, nodeMalformedPlatformID)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "//4232e3c7-d83c-d72b-758c-71d07a3d9310"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestPlatformLookupOnMalformedBlankPlatformIDField(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, nodeMalformedBlankPlatformID)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "other"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestPlatformLookupOnMalformedEmptyPlatformIDField(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, nodeMalformedEmptyPlatformID)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "other"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestPlatformLookupOnMalformedPartialPlatformIDField(t *testing.T) {
	t.Parallel()

	c := newTestCollectorForClusterWithNodes(t, nodeMalformedPartialPlatformID)
	got, err := c.Platform(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	want := "other"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

// newTestCollectorForClusterWithNodes returns a telemetry collector configured
// to simulate collecting data on a cluser with provided nodes.
func newTestCollectorForClusterWithNodes(t *testing.T, nodes ...runtime.Object) *telemetry.Collector {
	t.Helper()

	c, err := telemetry.NewCollector(
		telemetry.CollectorConfig{},
	)
	if err != nil {
		t.Fatal(err)
	}
	c.Config.K8sClientReader = newTestClientset(nodes...)
	return c
}

var (
	node1 = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "test-node-1",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{},
	}

	node2 = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "test-node-2",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{},
	}

	node3 = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "test-node-3",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{},
	}

	kubeNS = &apiCoreV1.Namespace{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name: "kube-system",
			UID:  "329766ff-5d78-4c9e-8736-7faad1f2e937",
		},
		Spec: apiCoreV1.NamespaceSpec{},
	}

	dummyKubeNS = &apiCoreV1.Namespace{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name: "kube-system",
			UID:  "",
		},
		Spec: apiCoreV1.NamespaceSpec{},
	}
)

// Cloud providers' nodes for testing ProviderID lookups.
var (
	nodeAWS = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "node",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{
			ProviderID: "aws:///eu-central-1a/i-088b4f07708408cc0",
		},
	}

	nodeAzure = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "node",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{
			ProviderID: "azure:///subscriptions/ba96ef31-4a42-40f5-8740-03f7e3c439eb/resourceGroups/mc_hibrid-weu_be3rr5ovr8ulf_westeurope/providers/Microsoft.Compute/virtualMachines/aks-pool1-27255451-0",
		},
	}

	nodeGCP = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "node",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{
			ProviderID: "gce://gcp-banzaidevgcp-nprd-38306/europe-north1-a/gke-vzf3z1vvleco9-pool1-7e48d363-8qz1",
		},
	}

	nodeKind = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "node",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{
			ProviderID: "kind://docker/local/local-control-plane",
		},
	}

	nodeVSphere = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "node",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{
			ProviderID: "vsphere://4232e3c7-d83c-d72b-758c-71d07a3d9310",
		},
	}

	nodeK3S = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "node",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{
			ProviderID: "k3s://ip-1.2.3.4",
		},
	}

	nodeIBMCloud = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "node",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{
			ProviderID: "ibmcloud://4232e3c7-d83c-d72b-758c-71d07a3d9310",
		},
	}

	nodeIBMPowerVS = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "node",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{
			ProviderID: "ibmpowervs://4232e3c7-d83c-d72b-758c-71d07a3d9310",
		},
	}

	nodeCloudStack = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "node",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{
			ProviderID: "cloudstack://4232e3c7-d83c-d72b-758c-71d07a3d9310",
		},
	}

	nodeOpenStack = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "node",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{
			ProviderID: "openstack://4232e3c7-d83c-d72b-758c-71d07a3d9310",
		},
	}

	nodeDigitalOcean = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "node",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{
			ProviderID: "digitalocean://4232e3c7-d83c-d72b-758c-71d07a3d9310",
		},
	}

	nodeEquinixMetal = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "node",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{
			ProviderID: "equinixmetal://4232e3c7-d83c-d72b-758c-71d07a3d9310",
		},
	}

	nodeAlibaba = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "node",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{
			ProviderID: "alicloud://4232e3c7-d83c-d72b-758c-71d07a3d9310",
		},
	}
)

// Nodes with missing or malformed PorviderID.
var (
	nodeMalformedPlatformID = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "node",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{
			ProviderID: "//4232e3c7-d83c-d72b-758c-71d07a3d9310",
		},
	}

	nodeMalformedPartialPlatformID = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "node",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{
			ProviderID: "://4232e3c7-d83c-d72b-758c-71d07a3d9310",
		},
	}

	nodeMalformedEmptyPlatformID = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "node",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{
			ProviderID: "",
		},
	}

	nodeMalformedBlankPlatformID = &apiCoreV1.Node{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      "node",
			Namespace: "default",
		},
		Spec: apiCoreV1.NodeSpec{
			ProviderID: " ",
		},
	}
)