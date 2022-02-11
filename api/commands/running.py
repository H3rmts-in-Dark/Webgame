import urllib.error

import click
from rich.console import Console

from GQL import Query
from Globlals import main

console = Console()


@main.command("running")
@click.option('-p', '--port', type=int, required=False, default=18266, show_default=True, help="specify port on server")
@click.option('-t', '--timeout', type=int, required=False, default=6, show_default=True, help="specify timeout for request")
@click.option('-v', '--verbose', 'verbosity', flag_value='verbose', is_flag=True, help="print more information about request")
@click.option('-q', '--quiet', 'verbosity', flag_value='quiet', is_flag=True, help="remove all printed information from request")
def _running(port: int, timeout: int, verbosity: str):
	data = Query("running.gql", port, timeout, verbosity == 'verbose', verbosity == 'quiet')
	if type(data) is urllib.error.URLError:
		console.print("Server not running", style="bold red")
	else:
		if verbosity == 'verbose':
			console.print(f"query:{data}", style="bold")
		console.print("Server running", style="bold green")
