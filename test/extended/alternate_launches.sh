#!/bin/bash
#
# This scripts starts the OpenShift server with a default configuration.
# The OpenShift Docker registry and router are installed.
# It will run all tests that are imported into test/extended.
source "$(dirname "${BASH_SOURCE}")/../../hack/lib/init.sh"

os::util::environment::setup_all_server_vars "test-extended-alternate-launches/"

export EXTENDED_TEST_PATH="${OS_ROOT}/test/extended"

function cleanup()
{
	out=$?
	pgrep -f "openshift" | xargs -r sudo kill
	cleanup_openshift

	# TODO(skuznets): un-hack this nonsense once traps are in a better state
	if [[ -n "${JUNIT_REPORT_OUTPUT:-}" ]]; then
		# get the jUnit output file into a workable state in case we crashed in
		# the middle of testing something
		os::test::junit::reconcile_output

		# check that we didn't mangle jUnit output
		os::test::junit::check_test_counters

		# use the junitreport tool to generate us a report
		os::util::ensure::built_binary_exists 'junitreport'

		cat "${JUNIT_REPORT_OUTPUT}" \
			| junitreport --type oscmd \
			--suites nested \
			--roots github.com/openshift/origin \
			--output "${ARTIFACT_DIR}/report.xml"
		cat "${ARTIFACT_DIR}/report.xml" | junitreport summarize
	fi

	os::log::info "Exiting"
	exit $out
}

trap "exit" INT TERM
trap "cleanup" EXIT

os::log::info "Starting server as distinct processes"
os::util::ensure::iptables_privileges_exist
os::start::configure_server

os::log::info "`openshift version`"
os::log::info "Server logs will be at:    ${LOG_DIR}/openshift.log"
os::log::info "Test artifacts will be in: ${ARTIFACT_DIR}"
os::log::info "Volumes dir is:            ${VOLUME_DIR}"
os::log::info "Config dir is:             ${SERVER_CONFIG_DIR}"
os::log::info "Using images:              ${USE_IMAGES}"
os::log::info "MasterIP is:               ${MASTER_ADDR}"

# Allow setting $JUNIT_REPORT to toggle output behavior
if [[ -n "${JUNIT_REPORT:-}" ]]; then
	export JUNIT_REPORT_OUTPUT="${LOG_DIR}/raw_test_output.log"
fi

mkdir -p ${LOG_DIR}

os::log::info "Scan of OpenShift related processes already up via ps -ef	| grep openshift : "
ps -ef | grep openshift

os::test::junit::declare_suite_start "extended/alternate_launches"

os::log::info "Starting etcdserver"
sudo env "PATH=${PATH}" OPENSHIFT_ON_PANIC=crash openshift start etcd \
 --config=${MASTER_CONFIG_DIR}/master-config.yaml \
 --loglevel=4 \
&>"${LOG_DIR}/os-etcdserver.log" &

os::log::info "Starting api server"
sudo env "PATH=${PATH}" OPENSHIFT_PROFILE=web OPENSHIFT_ON_PANIC=crash openshift start master api \
 --config=${MASTER_CONFIG_DIR}/master-config.yaml \
 --loglevel=4 \
&>"${LOG_DIR}/os-apiserver.log" &

os::cmd::try_until_text "oc get --raw /healthz --config='${MASTER_CONFIG_DIR}/admin.kubeconfig'" 'ok' $(( 80 * second )) 0.25
os::cmd::try_until_text "oc get --raw /healthz/ready --config='${MASTER_CONFIG_DIR}/admin.kubeconfig'" 'ok' $(( 80 * second )) 0.25
os::log::info "OpenShift API server up at: "
date

# test alternate node level launches
os::log::info "Testing alternate node configurations"

# proxy only
sudo env "PATH=${PATH}" TEST_CALL=1 OPENSHIFT_ON_PANIC=crash openshift start network --enable=proxy \
 --config=${NODE_CONFIG_DIR}/node-config.yaml \
 --loglevel=4 \
&>"${LOG_DIR}/os-network-1.log" &
os::cmd::try_until_text 'cat ${LOG_DIR}/os-network-1.log' 'syncProxyRules took'
pgrep -f "TEST_CALL=1" | xargs -r sudo kill
os::cmd::expect_success_and_text 'cat ${LOG_DIR}/os-network-1.log' 'Starting node networking'
os::cmd::expect_success_and_text 'cat ${LOG_DIR}/os-network-1.log' 'Started Kubernetes Proxy on'

# proxy only
sudo env "PATH=${PATH}" TEST_CALL=1 OPENSHIFT_ON_PANIC=crash openshift start node --enable=proxy \
 --config=${NODE_CONFIG_DIR}/node-config.yaml \
 --loglevel=4 \
&>"${LOG_DIR}/os-node-1.log" &
os::cmd::try_until_text 'cat ${LOG_DIR}/os-node-1.log' 'syncProxyRules took'
pgrep -f "TEST_CALL=1" | xargs -r sudo kill
os::cmd::expect_success_and_text 'cat ${LOG_DIR}/os-node-1.log' 'Starting node networking'
os::cmd::expect_success_and_text 'cat ${LOG_DIR}/os-node-1.log' 'Started Kubernetes Proxy on'

# plugins only
sudo env "PATH=${PATH}" TEST_CALL=1 OPENSHIFT_ON_PANIC=crash openshift start network --enable=plugins \
 --config=${NODE_CONFIG_DIR}/node-config.yaml \
 --loglevel=4 \
&>"${LOG_DIR}/os-network-2.log" &
os::cmd::try_until_text 'cat ${LOG_DIR}/os-network-2.log' 'Connecting to API server'
pgrep -f "TEST_CALL=1" | xargs -r sudo kill
os::cmd::expect_success_and_text 'cat ${LOG_DIR}/os-network-2.log' 'Starting node networking'
os::cmd::expect_success_and_not_text 'cat ${LOG_DIR}/os-network-2.log' 'Started Kubernetes Proxy on'

# plugins only
sudo env "PATH=${PATH}" TEST_CALL=1 OPENSHIFT_ON_PANIC=crash openshift start node --enable=plugins \
 --config=${NODE_CONFIG_DIR}/node-config.yaml \
 --loglevel=4 \
&>"${LOG_DIR}/os-node-2.log" &
os::cmd::try_until_text 'cat ${LOG_DIR}/os-node-2.log' 'Connecting to API server'
pgrep -f "TEST_CALL=1" | xargs -r sudo kill
os::cmd::expect_success_and_text 'cat ${LOG_DIR}/os-node-2.log' 'Starting node networking'
os::cmd::expect_success_and_not_text 'cat ${LOG_DIR}/os-node-2.log' 'Started Kubernetes Proxy on'

# kubelet only
sudo env "PATH=${PATH}" TEST_CALL=1 OPENSHIFT_ON_PANIC=crash openshift start node --enable=kubelet \
 --config=${NODE_CONFIG_DIR}/node-config.yaml \
 --loglevel=4 \
&>"${LOG_DIR}/os-node-3.log" &
os::cmd::try_until_text 'cat ${LOG_DIR}/os-node-3.log' 'Started kubelet'
pgrep -f "TEST_CALL=1" | xargs -r sudo kill
os::cmd::expect_success_and_text 'cat ${LOG_DIR}/os-node-3.log' 'Starting node'
os::cmd::expect_success_and_not_text 'cat ${LOG_DIR}/os-node-3.log' 'Starting node networking'
os::cmd::expect_success_and_not_text 'cat ${LOG_DIR}/os-node-3.log' 'Started Kubernetes Proxy on'


os::log::info "Starting controllers"
sudo env "PATH=${PATH}"  OPENSHIFT_ON_PANIC=crash openshift start master controllers \
 --config=${MASTER_CONFIG_DIR}/master-config.yaml \
 --loglevel=4 \
&>"${LOG_DIR}/os-controllers.log" &

os::log::info "Starting node"
sudo env "PATH=${PATH}"  OPENSHIFT_ON_PANIC=crash openshift start node \
 --config=${NODE_CONFIG_DIR}/node-config.yaml \
 --loglevel=4 \
&>"${LOG_DIR}/os-node.log" &
export OS_PID=$!

os::log::info "OpenShift server start at: "
date

os::cmd::try_until_text "oc get --raw ${KUBELET_SCHEME}://${KUBELET_HOST}:${KUBELET_PORT}/healthz --config='${MASTER_CONFIG_DIR}/admin.kubeconfig'" 'ok' minute 0.5
os::cmd::try_until_success "oc get --raw /api/v1/nodes/${KUBELET_HOST} --config='${MASTER_CONFIG_DIR}/admin.kubeconfig'" $(( 80 * second )) 0.25
os::log::info "OpenShift node health checks done at: "
date

os::test::junit::declare_suite_end

# set our default KUBECONFIG location
export KUBECONFIG="${ADMIN_KUBECONFIG}"

${OS_ROOT}/test/end-to-end/core.sh
