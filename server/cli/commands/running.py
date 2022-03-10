import urllib.error

import click
from click import Context

from GQL import genQuery
from Globlals import main, pr, log, Container

gql = """
query {
    ping {
        uptime
    }
}
"""


@main.command("running", help="Checks if a goServer is running at host:port")
@click.pass_context
def _running(ctx: Context):
	container: Container = ctx.obj
	data = genQuery(container, gql)
	if type(data) is urllib.error.URLError:
		if data.reason.strerror == "Connection refused":
			pr(f"error:{data.reason}", style="red")
			pr("Server not running", style="bold red")
		else:
			log(container.verbosity, f"error:{data.reason}", style="red")
			pr("Error connecting to Server", style="bold red")
	else:
		log(container.verbosity, f"query:{data}")
		pr("Server running", style="bold green")
