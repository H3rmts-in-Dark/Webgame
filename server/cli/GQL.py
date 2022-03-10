# https://github.com/profusion/sgqlc

import urllib.error
from typing import Union

from sgqlc.endpoint.http import HTTPEndpoint

from Globlals import Logger, log, Verbosity, Container


def Query(query: str, host: str, port: int, https: bool, timeout: int, verbose: Verbosity) -> Union[
	dict, urllib.error.URLError]:
	endpoint = HTTPEndpoint("{2}://{0}:{1}/query".format(host, port, "https" if https else "http"), timeout=timeout)
	log(verbose, f"sending to {endpoint.url} with timeout {endpoint.timeout}")
	endpoint.logger = Logger(verbose)
	try:
		return endpoint(query, timeout=timeout)
	except urllib.error.URLError as ex:
		return ex


def genQuery(container: Container, gql):
	return Query(gql, container.host, container.port, container.https, container.timeout, container.verbosity)
