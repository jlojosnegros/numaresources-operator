# NUMA Resources Operator test suite

The NUMA resources operator e2e suite can *optionally* setup/teardown the numaresources stack, but the suite expects a pre-configured kubeletconfig.
Please find [here](https://raw.githubusercontent.com/openshift-kni/numaresources-operator/main/doc/examples/kubeletconfig.yaml) an example of recommended kubeletconfig.

The e2e suite can set a default `kubeletconfig`, but this is not recommended. The recommended flow is to pre-configure the cluster with a `kubeletconfig` object.
Should you decide to use the default `kubeletconfig`, please omit the `-e E2E_NROP_INSTALL_SKIP_KC=true` from all the `podman` command lines below.

The e2e suite assumes the cluster has the numaresources operator installed, but with no configuration. To install the numaresources operator, you can use the vehicle which best suits your use case (OLM, `make deploy`...).

In addition to kubelet config, refer to [doc/examples](doc/examples) directory for example specification files.  NUMAResourcesOperator CR and NUMAResourcesScheduler CR must be deployed in the cluster. These CRs take care of deployment of Resouce Topology Exporter daemon (on nodes that are targetted by the machine config pool selector) and NUMA-aware scheduler plugin as a secondary scheduler plugin respective.

## Running from the source tree

This is the easiest way, recommended for developers and people which want to consume the latest and greatest source tree.
From the NUMA Resources Operator source tree:

```bash
export KUBECONFIG=...
./hack/run-test-serial-e2e.sh
```

To get all the options supported by the run script:

```bash
./hack/run-test-serial-e2e.sh --help
```

## Running from container images

[Pre-built container images](https://quay.io/repository/openshift-kni/numaresources-operator-tests) are available which contain the full testsuite, to enable containerized flows.

To actually run the tests and the tests only, assuming a pre-configured numaresources stack

```bash
podman run -ti \
 -v $KUBECONFIG:/kubeconfig:z \
 -e KUBECONFIG=/kubeconfig
 -e E2E_NROP_INSTALL_SKIP_KC=true \
 quay.io/openshift-kni/numaresources-operator-tests:4.14.999-snapshot
```

To setup the stack from scratch and then run the tests (you may want to do that with *ephemeral* CI clusters)

```bash
podman run -ti \
 -v $KUBECONFIG:/kubeconfig:z \
 -e KUBECONFIG=/kubeconfig \
 -e E2E_NROP_INSTALL_SKIP_KC=true \
 quay.io/openshift-kni/numaresources-operator-tests:4.14.999-snapshot \
 --setup
```

To setup the stack, run the tests and then restore the pristine cluster state:

```bash
podman run -ti \
 -v $KUBECONFIG:/kubeconfig:z \
 -e KUBECONFIG=/kubeconfig \
 -e E2E_NROP_INSTALL_SKIP_KC=true \
 quay.io/openshift-kni/numaresources-operator-tests:4.14.999-snapshot \
 --setup \
 --teardown
```

## avoiding dependencies on other images

The E2E suite depends on few extra images. These images are very stable, lightweight and little concern most of times:

- `quay.io/openshift-kni/numacell-device-plugin:test-ci`
- `quay.io/openshift-kni/pause:test-ci`

However, in some cases it may be unpractical to depend on third party images.
The E2E test image can act as replacement for all its dependencies, providing either the same code or replacements suitables for its use case.
To replace the dependencies, you need to set up some environment variables:

```bash
export E2E_IMAGE_URL=quay.io/openshift-kni/numaresources-operator-tests:4.14.999-snapshot
podman run -ti \
 -v $KUBECONFIG:/kubeconfig:z \
 -e KUBECONFIG=/kubeconfig \
 -e E2E_NROP_INSTALL_SKIP_KC=true \
 -e E2E_NUMACELL_DEVICE_PLUGIN_URL=${E2E_IMAGE_URL} \
 -e E2E_PAUSE_IMAGE_URL=${E2E_IMAGE_URL} \
 ${E2E_IMAGE_URL}
```

## running the tests using cnf-tests

The [CNF tests](https://github.com/openshift-kni/cnf-features-deploy/blob/master/cnf-tests/README.md) [container images](https://quay.io/repository/openshift-kni/cnf-tests) includes the E2E suite.
While the primary source for pre-built test container image is the [numaresources-operator-tests](https://quay.io/repository/openshift-kni/numaresources-operator-tests), the CNF tests integration
will be updated shortly after. **Running the testsuite through CNF tests is fully supported**.
To run the suite using the CNF tests image, you can run

```bash
export CNF_TESTS_URL="quay.io/openshift-kni/cnf-tests:4.14.0"
podman run -ti \
 -v $KUBECONFIG:/kubeconfig:z \
 -e KUBECONFIG=/kubeconfig \
 -e CLEAN_PERFORMANCE_PROFILE=false \
 -e E2E_NROP_INSTALL_SKIP_KC=true \
 -e E2E_NUMACELL_DEVICE_PLUGIN_URL=${CNF_TESTS_URL} \
 -e E2E_PAUSE_IMAGE_URL=${CNF_TESTS_URL} \
 ${CNF_TESTS_URL} \
 /usr/bin/test-run.sh \
 -ginkgo.focus="numaresources"
```

## skipping reboot-requiring tests

Some E2E tests require to reboot one or more worker node. This is intrinsically fragile and slow, and you may want to avoid to do this in your tier-1 runs.
To do so, you can run

```bash
export E2E_IMAGE_URL=quay.io/openshift-kni/numaresources-operator-tests:4.14.999-snapshot
podman run -ti \
 -v $KUBECONFIG:/kubeconfig:z \
 -e KUBECONFIG=/kubeconfig \
 -e E2E_NROP_INSTALL_SKIP_KC=true \
 -e E2E_NUMACELL_DEVICE_PLUGIN_URL=${E2E_IMAGE_URL} \
 -e E2E_PAUSE_IMAGE_URL=${E2E_IMAGE_URL} \
 ${E2E_IMAGE_URL}
 --skip '.*reboot_required.*'
```

or, with CNF tests:

```bash
export CNF_TESTS_URL="quay.io/openshift-kni/cnf-tests:4.14.0"
podman run -ti \
 -v $KUBECONFIG:/kubeconfig:z \
 -e KUBECONFIG=/kubeconfig \
 -e CLEAN_PERFORMANCE_PROFILE=false \
 -e E2E_NROP_INSTALL_SKIP_KC=true \
 -e E2E_NUMACELL_DEVICE_PLUGIN_URL=${CNF_TESTS_URL} \
 -e E2E_PAUSE_IMAGE_URL=${CNF_TESTS_URL} \
 ${CNF_TESTS_URL} \
 /usr/bin/test-run.sh \
 -ginkgo.skip='.*reboot_required.*' \
 -ginkgo.focus="numaresources"
```
