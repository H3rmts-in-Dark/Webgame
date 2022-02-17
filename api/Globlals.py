from typing import Literal, Union, Optional

import click
import rich
from rich import console

c = console.Console()


@click.group()
def main():
    """
    CLI to interact with golang server
    """


def pr(verbosity: Union[Literal["verbose", "quiet"], None], msg: str, *args, style: Optional[Union[str, rich.console.Style]] = None):
    """
    prints message with style provided

    if set to None it just prints
    if set to quiet it just prints
    """
    c.print(msg % args, style=style)


def log(verbosity: Union[Literal["verbose", "quiet"], None], msg: str, *args, style: Optional[Union[str, rich.console.Style]] = None):
    """
    prints message with grey color

    only logs if verbosity is set to "verbose" or none
    if set to verbose it prints debug infos
    if set to None it just prints
    if set to quiet it skips
    """
    if verbosity == "verbose":
        c.log(msg % args, style=console.Style(color="rgb(85,85,85)") if style is None else style)
    elif verbosity is None:
        c.print(msg % args, style=console.Style(color="rgb(85,85,85)") if style is None else style)
