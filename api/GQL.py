import logging
import urllib.error
from typing import Union

from sgqlc.endpoint.http import HTTPEndpoint


# https://github.com/profusion/sgqlc

def Query(name: str, host: str, port: int, https: bool, timeout: int, verbose: bool, quiet: bool) -> Union[dict, urllib.error.URLError]:
	with open("commands/{0}".format(name)) as f:
		endpoint = HTTPEndpoint("{2}://{0}:{1}/query".format(host, port, "https" if https else "http"), timeout=timeout)
		if quiet:
			endpoint.logger.disabled = True
		elif verbose:
			logging.basicConfig(level=logging.DEBUG)
		
		try:
			return endpoint(f.read(), timeout=timeout)
		except urllib.error.URLError as ex:
			return ex
