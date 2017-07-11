#!/usr/bin/env python
import re
import json
from subprocess import Popen, PIPE
from prometheus_client import start_http_server, Summary, Gauge
from optparse import OptionParser

parser = OptionParser()

parser.add_option("-l", "--listen", dest="listen", help="ip/hostname on which server will listen", action="store", default="0.0.0.0")
parser.add_option("-p", "--port", dest="port", help="port on which server will listen", action="store", default="8000")

(options, args) = parser.parse_args()

# Create a metric to track time spent and requests made.
REQUEST_TIME = Summary('request_processing_seconds', 'Time spent processing request')

namespace_limits = ["hard", "used"]

gauges = {}

prefix_namespace_limit = "k8s_namespace"

environment = "dev"

gauges["%s_info" % (prefix_namespace_limit)] = Gauge(
    "%s_info" % (prefix_namespace_limit),
    '%s_info' % (prefix_namespace_limit),
    ['environment', 'namespace']
)

for item in namespace_limits:
    gauges["%s_quota_%s_limits_cpu" % (prefix_namespace_limit, item)] = Gauge(
        "%s_quota_%s_limits_cpu" % (prefix_namespace_limit, item),
        '%s_quota_%s_limits_cpu' % (prefix_namespace_limit, item),
        ['environment', 'namespace']
    )
    gauges["%s_quota_%s_limits_memory" % (prefix_namespace_limit, item)] = Gauge(
        "%s_quota_%s_limits_memory" % (prefix_namespace_limit, item),
        '%s_quota_%s_limits_memory' % (prefix_namespace_limit, item),
        ['environment', 'namespace']
    )
    gauges["%s_quota_%s_pods" % (prefix_namespace_limit, item)] = Gauge(
        "%s_quota_%s_pods" % (prefix_namespace_limit, item),
        '%s_quota_%s_pods' % (prefix_namespace_limit, item),
        ['environment', 'namespace']
    )
    gauges["%s_quota_%s_requests_cpu" % (prefix_namespace_limit, item)] = Gauge(
        "%s_quota_%s_requests_cpu" % (prefix_namespace_limit, item),
        '%s_quota_%s_requests_cpu' % (prefix_namespace_limit, item),
        ['environment', 'namespace']
    )
    gauges["%s_quota_%s_requests_memory" % (prefix_namespace_limit, item)] = Gauge(
        "%s_quota_%s_requests_memory" % (prefix_namespace_limit, item),
        '%s_quota_%s_requests_memory' % (prefix_namespace_limit, item),
        ['environment', 'namespace']
    )

gauges["%s_limits_cpu" % (prefix_namespace_limit)] = Gauge(
    "%s_limits_cpu" % (prefix_namespace_limit),
    '%s_limits_cpu' % (prefix_namespace_limit),
    ['environment', 'namespace', 'id', 'config']
)

gauges["%s_limits_memory" % (prefix_namespace_limit)] = Gauge(
    "%s_limits_memory" % (prefix_namespace_limit),
    '%s_limits_memory' % (prefix_namespace_limit),
    ['environment', 'namespace', 'id', 'type']
)


def calculate_metric(value):
    """Calculate metric."""
    tmp_value = re.match("^(\d+)([\D]+)$", value)

    multiplier = 1
    if tmp_value and tmp_value.groups():
        if "Mi" == tmp_value.group(2):
            multiplier = 1024 * 1024
        elif "Gi" == tmp_value.group(2):
            multiplier = 1024 * 1024 * 1024
        elif "m" == tmp_value.group(2):
            multiplier = 0.001

        tmp_value = float(tmp_value.group(1))
    else:
        tmp_value = value

    return tmp_value * multiplier


def get_quotas():
    """Get quotas."""
    command = [
        "kubectl",
        "--all-namespaces=true",
        "get",
        "quota",
        "-o",
        "json"
    ]

    p = Popen(command, stdin=PIPE, stdout=PIPE, stderr=PIPE)

    (command_output, command_error) = p.communicate()

    command_output = command_output.strip()
    command_error = command_error.strip()

    quotas = json.loads(command_output)

    for quota in quotas["items"]:

        for item in namespace_limits:
            gauges["%s_quota_%s_limits_cpu" % (prefix_namespace_limit, item)].labels(environment, quota["metadata"]["name"]).set(calculate_metric(quota["status"][item]["limits.cpu"]))
            gauges["%s_quota_%s_limits_memory" % (prefix_namespace_limit, item)].labels(environment, quota["metadata"]["name"]).set(calculate_metric(quota["status"][item]["limits.memory"]))
            gauges["%s_quota_%s_requests_cpu" % (prefix_namespace_limit, item)].labels(environment, quota["metadata"]["name"]).set(calculate_metric(quota["status"][item]["requests.cpu"]))
            gauges["%s_quota_%s_requests_memory" % (prefix_namespace_limit, item)].labels(environment, quota["metadata"]["name"]).set(calculate_metric(quota["status"][item]["requests.memory"]))
            gauges["%s_quota_%s_pods" % (prefix_namespace_limit, item)].labels(environment, quota["metadata"]["name"]).set(int(quota["status"][item]["pods"]))


def get_limits():
    """Get limits."""
    command = [
        "kubectl",
        "--all-namespaces=true",
        "get",
        "limits",
        "-o",
        "json"
    ]

    p = Popen(command, stdin=PIPE, stdout=PIPE, stderr=PIPE)

    (command_output, command_error) = p.communicate()

    command_output = command_output.strip()
    command_error = command_error.strip()

    limits = json.loads(command_output)

    for limit in limits["items"]:
        i = 0
        for item in limit["spec"]["limits"]:
            i += 1
            for limit_item in item:
                if type(item[limit_item]) is dict:
                    gauges["%s_limits_cpu" % (prefix_namespace_limit)].labels(environment, limit["metadata"]["name"], i, limit_item).set(calculate_metric(item[limit_item]["cpu"]))
                    gauges["%s_limits_memory" % (prefix_namespace_limit)].labels(environment, limit["metadata"]["name"], i, limit_item).set(calculate_metric(item[limit_item]["memory"]))


def get_namespaces():
    """Get limits."""
    command = [
        "kubectl",
        "get",
        "namespaces",
        "-o",
        "json"
    ]

    p = Popen(command, stdin=PIPE, stdout=PIPE, stderr=PIPE)

    (command_output, command_error) = p.communicate()

    command_output = command_output.strip()
    command_error = command_error.strip()

    namespaces = json.loads(command_output)

    for namespace in namespaces["items"]:
        gauges["%s_info" % (prefix_namespace_limit)].labels(environment, namespace["metadata"]["name"]).set(1)


# Decorate function with metric.
@REQUEST_TIME.time()
def process_request():
    """Process request."""
    get_namespaces()
    get_quotas()
    get_limits()

if __name__ == '__main__':
    # Start up the server to expose the metrics.
    start_http_server(int(options.port), addr=options.listen)
    while True:
        process_request()
