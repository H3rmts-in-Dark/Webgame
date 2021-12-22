from pyfiglet import Figlet

f = Figlet(font='slant')
print(f.renderText('Goserver'))

import click
import requests

__author__ = "Oyetoke Toby"


@click.group()
def main():
	"""
	Simple CLI for querying books on Google Books by Oyetoke Toby
	"""
	pass


@main.command()
@click.argument('query')
def search(query):
	"""This search and return results corresponding to the given query from Google Books"""
	url_format = 'https://www.googleapis.com/books/v1/volumes'
	query = "+".join(query.split())
	
	query_params = {
		'q': query
	}
	
	response = requests.get(url_format, params=query_params)
	
	click.echo(response.json()['items'])


@main.command()
@click.argument('id')
def get(id):
	"""This return a particular book from the given id on Google Books"""
	url_format = 'https://www.googleapis.com/books/v1/volumes/{}'
	click.echo(id)
	
	response = requests.get(url_format.format(id))
	
	click.echo(response.json())


if __name__ == "__main__":
	main()
	input()
