import click
from PyInquirer import style_from_dict, Token, prompt

from Globlals import main
from commands import running

_ = running


# https://zetcode.com/python/click/

# https://click.palletsprojects.com/en/7.x/options/


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
