import logging
import urllib.error
from typing import Union

from sgqlc.endpoint.http import HTTPEndpoint


# https://github.com/profusion/sgqlc

def Query(name: str, port: int, timeout: int, verbose: bool, quiet: bool) -> Union[dict, urllib.error.URLError]:
	with open("commands/{0}".format(name)) as f:
		endpoint = HTTPEndpoint("http://localhost:{0}/query".format(port), timeout=timeout)
		if quiet:
			endpoint.logger.disabled = True
		elif verbose:
			logging.basicConfig(level=logging.DEBUG)
		
		try:
			response = endpoint(f.read())
			return response
		except urllib.error.URLError as ex:
			return ex
