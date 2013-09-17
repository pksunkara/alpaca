#!/usr/bin/env python
import sys

try:
	from setuptools import setup
except ImportError:
	from distutils.core import setup

extra = {}

requires = [
	'pystache==0.5.3',
	'jsonschema==2.0.0',
	'simplejson==3.3.0'
]

if sys.version_info[:2] == (2, 4):
	# For python2.4 we use simplejson 2.0.9
	requires[0] = 'simplejson<2.1.0'

if sys.version_info[:2] == (2, 6):
	# For python2.6 we have to require argparse since it was not in stdlib until 2.7.
	requires.append('argparse>=1.1')

if sys.version_info >= (3,):
	# Convert python2.x to python3.x
	extra['use_2to3'] = True

setup(
	name = 'alpaca-cli',
	version = '1.0',
	description = 'Api Libraries Powered And Created by Alpaca',
	long_description = open('README.md').read(),
	author = 'Pavan Kumar Sunkara',
	author_email = 'pavan.sss1991@gmail.com',
	url = 'http://github.com/pksunkara/alpaca',
	entry_points={
		'console_scripts': [
			'alpaca = alpaca.__main__:main',
		],
	},
	license = 'MIT',
	keywords = 'api libraries automate rest',
	packages = ['alpaca'],
	install_requires = requires,
	include_package_data = True,
	zip_safe = True,
	classifiers = (
		'Development Status :: 2 - Pre-Alpha',
		'Environment :: Console',
		'Intended Audience :: Developers',
		'License :: OSI Approved :: MIT License',
		'Programming Language :: Python',
		'Programming Language :: Python :: 2.6',
		'Programming Language :: Python :: 2.7',
		'Programming Language :: Python :: 3.1',
		'Programming Language :: Python :: 3.2',
		'Programming Language :: Python :: 3.3',
		'Topic :: Software Development :: Libraries'
	),
	**extra
)
