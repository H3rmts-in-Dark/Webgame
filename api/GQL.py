import urllib.error

from sgqlc.endpoint.http import HTTPEndpoint

# https://github.com/profusion/sgqlc 

def Query(name: str, port: int, timeout: int):
	with open("commands/{0}".format(name)) as f:
		endpoint = HTTPEndpoint("http://localhost:{0}/query".format(port), timeout=timeout)
		try:
			return type(endpoint(f.read()))
		except urllib.error.URLError as ex:
			return ex
