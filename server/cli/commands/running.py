import urllib.error

import click
from click import Context

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
	data = container.genQuery(gql)
	if type(data) is urllib.error.URLError:
		if data.reason.strerror == "Connection refused":
			pr(container.verbosity, f"error:{data.reason}", style="red")
			pr(container.verbosity, "Server not running", style="bold red")
		else:
			log(container.verbosity, f"error:{data.reason}", style="red")
			pr(container.verbosity, "Error connecting to Server", style="bold red")
	else:
		log(container.verbosity, f"query:{data}")
		pr(container.verbosity, "Server running", style="bold green")
