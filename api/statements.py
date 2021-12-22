from dataclasses import dataclass
from typing import List
from xml.etree import ElementTree

from lxml import etree


@dataclass
class Argument:
	name: str
	type: str


@dataclass
class Help:
	name: str
	type: str


@dataclass
class Command:
	name: str
	arguments: List[Argument]
	help: Help


def readcommands():  # read a certain command from xml structure
	with open("commands.xml") as file:
		root2: ElementTree.Element = etree.XML("".join(file.readlines()))
		for child in root2:
			print(child)
			for child2 in child:
				print("  ", child2)
				for child3 in child2:
					print("    ", child3)
					for child4 in child3:
						print("      ", child4)


def switch(command):
	pass
