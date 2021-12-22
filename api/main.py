import sys

from gql import Client
from gql.transport.aiohttp import AIOHTTPTransport

import QLtest
import statements

var = sys.argv
print(var)

client = Client(transport=AIOHTTPTransport(url="http://localhost:18266/query"), fetch_schema_from_transport=True)
QLtest.test(client)
statements.readcommands()
statements.switch("test")

# https://codeburst.io/building-beautiful-command-line-interfaces-with-python-26c7e1bb54df
