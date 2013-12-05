import os
import sys

from setuptools import setup

setup(
	name='{{.Pkg.package}}',
	version='{{.Pkg.version}}',
	description='{{if .Pkg.official}}Official {{end}}{{.Pkg.name}} API library client for python',
	author='{{.Pkg.author.name}}',
	author_email='{{.Pkg.author.email}}',
	url='{{.Pkg.url}}',
	license='{{.Pkg.license}}',
	install_requires=[],
	packages=[
		'{{call .Fnc.underscore .Pkg.name}}'
	],
	classifiers=[
		'Intended Audience :: Developers',
		'License :: OSI Approved :: {{if .Pkg.python}}{{or .Pkg.python.license .Pkg.license}}{{end}}',
		'Operating System :: OS Independent',
		'Topic :: Software Development :: Libraries :: Python Modules',
	],
	use_2to3 = True
)
