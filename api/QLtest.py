from gql import Client, gql
from gql.transport.exceptions import TransportQueryError


def test(client: Client):
	query = gql(
		"""
		mutation {
			reloadSites(validation: "APICode") {
				ok
				info
			}
		}
		"""
	)
	
	# Execute the query on the transport
	try:
		result = client.execute(query)
		print(result)
	except TransportQueryError as tqe:
		print(tqe)
