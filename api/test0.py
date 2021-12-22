import argparse

parser = argparse.ArgumentParser()
parser.add_argument(dest='argument1', help="This is the first argument")

parser.add_argument(dest='argument1', type=str, help="A string argument")
parser.add_argument(dest='argument2', type=int, help="An integer argument")
parser.add_argument(dest='argument3', type=float, help="A float argument")

# Validate that the input is in specified list
parser.add_argument(dest='argument4', choices=['red', 'green', 'blue'])

# Accept multiple inputs for an argument, returned as a list
# Will be of type string, unless specified
parser.add_argument(dest='argument5', nargs=2, type=int)

# Optional positional argument (length 0 or 1)
parser.add_argument(dest='argument6', nargs='?')

# Boolean flag (does not accept input data), with default value
parser.add_argument('-a1', action="store_true", default=False)

# Cast input to integer, with a default value
parser.add_argument('-a2', type=int, default=0)

# Provide long form name as well (maps to 'argument3' not 'a3')
parser.add_argument('-a3', '--argument3', type=str)

# Make argument mandatory
parser.add_argument('-a4', required=True)

# Retur the input via different parameter name
parser.add_argument('-a5', '--argument5', dest='my_argument')

group = parser.add_mutually_exclusive_group()
group.add_argument('--arg1', action='store_true')
group.add_argument('--arg2', action='store_false')


def single_word(string):
	# Check input does not contain spaces
	if ' ' in string:
		msg = f'\"{string}\" is not a single word'
		raise argparse.ArgumentTypeError(msg)
	return string


parser.add_argument('argument1', type=single_word)

args = parser.parse_args()
print(args.a1)
print(args.a2)
print(args.argument3)
print(args.a4)
print(args.my_argument)
