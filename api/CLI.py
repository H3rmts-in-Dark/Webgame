import click
import requests
from PyInquirer import style_from_dict, Token, prompt


# https://zetcode.com/python/click/


@click.group()
def main():
	"""
	CLI to interact with golang server
	"""


@main.command("gethello")
@click.argument('age', type=int)
@click.argument('name', default='guest')
@click.option('--shout/--no-shout', default=False, help="This help message")
def hello(name, age, shout):
	"""
	Test hello
	"""
	click.echo(f'{name} is {age} years old {shout}')


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


@main.command("dial")
def dialog():
	"""Dialog"""
	questions = [
		{
			'type': 'confirm',
			'name': 'toBeDelivered',
			'message': 'Is this for delivery?',
			'default': False
		},
		{
			'type': 'input',
			'name': 'phone',
			'message': 'What\'s your phone number?',
		},
		{
			'type': 'list',
			'name': 'size',
			'message': 'What size do you need?',
			'choices': ['Large', 'Medium', 'Small'],
			'filter': lambda val: val.lower()
		},
		{
			'type': 'input',
			'name': 'quantity',
			'message': 'How many do you need?',
			'filter': lambda val: int(val)
		},
		{
			'type': 'input',
			'name': 'comments',
			'message': 'Any comments on your purchase experience?',
			'default': 'Nope, all good!'
		}
	]
	style = style_from_dict({
		Token.QuestionMark: '#E91E63 bold',
		Token.Selected: '#673AB7 bold',
		Token.Instruction: '#F91EB3',
		Token.Answer: '#2196f3 bold',
		Token.Question: '#5012f3',
	})
	answers = prompt(questions, style=style)
	click.echo(answers)

if __name__ == "__main__":
	main()

# https://codeburst.io/building-beautiful-command-line-interfaces-with-python-26c7e1bb54df
