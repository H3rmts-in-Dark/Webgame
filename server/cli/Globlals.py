import dataclasses
import logging
from typing import Optional
from typing import Union, Literal

import click
import rich
from click import Context
from rich import console

c = console.Console()

Verbosity = Union[Literal["verbose", "quiet", "debug"], None]


@dataclasses.dataclass
class Container:
	host: str
	port: int
	timeout: int
	verbosity: Verbosity
	https: bool


@click.group()
@click.option('-h', '--host', type=str, required=False, default="localhost", show_default=True, help="specify host of server")
@click.option('-p', '--port', type=int, required=False, default=18266, show_default=True, help="specify port on server")
@click.option('-t', '--timeout', type=int, required=False, default=6, show_default=True, help="specify timeout for request")
@click.option('-v', '--verbose', 'verbosity', flag_value='verbose', is_flag=True, help="print more information about request")
@click.option('-q', '--quiet', 'verbosity', flag_value='quiet', is_flag=True, help="remove all printed information from request")
@click.option('-d', '--debug', 'verbosity', flag_value='debug', is_flag=True, help="prints local variables and every debug information")
@click.option('--https/--http', type=bool, default=True, show_default=True, help="Use https or http as connection")
@click.pass_context
def main(ctx: Context, host: str, port: int, timeout: int, verbosity: Verbosity, https: bool):
	"""
	CLI to interact with goServer
	"""
	ctx.obj = Container(host, port, timeout, verbosity, https)


def pr(msg: str, *args, style: Optional[Union[str, rich.console.Style]] = None):
	"""
	prints message with style provided

	if set to None it just prints
	if set to quiet it just prints
	"""
	c.print(msg % args, style=style)


def log(verbosity: Union[Literal["verbose", "quiet"], None], msg: str, *args, style: Optional[Union[str, rich.console.Style]] = None,
        skip=0):
	"""
	prints message with grey color

	only logs if verbosity is set to "verbose" or none
	if set to debug it prints debug infos and local variables
	if set to verbose it prints debug infos
	if set to None it just prints
	if set to quiet it skips
	"""
	if verbosity == "verbose" or verbosity == "debug":
		c.log(msg % args, style=console.Style(color="rgb(85,85,85)") if style is None else style, log_locals=verbosity == "debug",
		      _stack_offset=(2 + skip))
	elif verbosity is None:
		c.print(msg % args, style=console.Style(color="rgb(85,85,85)") if style is None else style)


def deb(offset: int = 1):
	"""
	
	"""
	c.log(log_locals=True, _stack_offset=1 + offset)


class Logger(logging.Logger):
	def __init__(self, verbosity: Union[Literal["verbose", "quiet"], None]):
		super().__init__('GQL Logger')
		self.verbosity = verbosity
	
	def debug(self, msg: str, *args, **kwargs) -> None:
		if self.verbosity == "verbose" or self.verbosity == "debug":
			log(self.verbosity, msg % args, skip=1)
	
	def exception(self, msg: str, *args, **kwargs) -> None:
		log(self.verbosity, msg % args, skip=1)
	
	def info(self, msg: str, *args, **kwargs) -> None:
		log(self.verbosity, msg % args, skip=1)
	
	def error(self, msg: str, *args, **kwargs) -> None:
		log(self.verbosity, msg % args, skip=1)
	
	def warning(self, msg: str, *args, **kwargs) -> None:
		log(self.verbosity, msg % args, skip=1)
	
	def critical(self, msg: str, *args, **kwargs) -> None:
		log(self.verbosity, msg % args, skip=1)
	
	def log(self, msg: str, *args, **kwargs) -> None:
		if self.verbosity != "clean":
			log(self.verbosity, msg % args, skip=1)
