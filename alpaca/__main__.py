#!/usr/bin/env python

"""

The main entry point. Invoke as 'alpaca' or 'python -m alpaca'

"""

import sys

from .core import main

if __name__ == '__main__':
	sys.exit(main())
