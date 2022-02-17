import logging
import urllib.error
from typing import Union, Literal

import rich
from rich.console import Console
from sgqlc.endpoint.http import HTTPEndpoint

from Globlals import log

console = Console()


class Logger(logging.Logger):
    def __init__(self, verbosity: Union[Literal["verbose", "quiet"], None]):
        super().__init__('GQL Logger')
        self.verbosity = verbosity

    def debug(self, msg: str, *args, **kwargs) -> None:
        if self.verbosity == "verbose":
            log(self.verbosity, msg % args)

    def exception(self, msg: str, *args, **kwargs) -> None:
        log(self.verbosity, msg % args)

    def info(self, msg: str, *args, **kwargs) -> None:
        log(self.verbosity, msg % args)

    def error(self, msg: str, *args, **kwargs) -> None:
        log(self.verbosity, msg % args)

    def warning(self, msg: str, *args, **kwargs) -> None:
        log(self.verbosity, msg % args)

    def critical(self, msg: str, *args, **kwargs) -> None:
        log(self.verbosity, msg % args)

    def log(self, msg: str, *args, **kwargs) -> None:
        if self.verbosity != "clean":
            log(self.verbosity, msg % args)


# https://github.com/profusion/sgqlc
def Query(
        name: str, host: str, port: int,
        https: bool, timeout: int, verbose: Union[Literal["verbose", "quiet"], None]
) -> Union[dict, urllib.error.URLError]:
    with open("commands/{0}".format(name)) as f:
        endpoint = HTTPEndpoint("{2}://{0}:{1}/query".format(host, port, "https" if https else "http"), timeout=timeout)
        log(verbose, f"sending to {endpoint.url} with timeout {endpoint.timeout}")
        endpoint.logger = Logger(verbose)
        try:
            return endpoint(f.read(), timeout=timeout)
        except urllib.error.URLError as ex:
            return ex
