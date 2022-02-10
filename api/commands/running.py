import click
from rich import Console

from CLI import main
from GQL import Query

console = Console()

@main.command("running")
@click.argument('port', type=int, required=False, default=18266)
@click.argument('timeout', type=int, required=False, default=12)
@click.option('-v')
def _running(port):
	print(Query("running.gql", port, 10))
	try:
		console.print("query", style="bold")
	except aiohttp.ClientConnectorError as ce:
		console.print("Server not running", style="bold red")
	except TransportQueryError as tqe:
		click.echo(tqe, err=True, color=True)
