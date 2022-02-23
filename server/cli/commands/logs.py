import click
from click import Context

from Globlals import main, Container


@main.group('logs', help='test', invoke_without_command=True)
@click.option('--style', type=str, required=False, help="style")
@click.pass_context
def _logs(ctx: Context):
	container: Container = ctx.obj
	if ctx.invoked_subcommand is None:
		pass
	pass


@_logs.command("limit", help='returns limited count of logs')
@click.argument('limit', type=int, required=False, default=500)
@click.pass_context
def _log_limit(ctx: Context, limit: int):
	container: Container = ctx.obj
	pass
