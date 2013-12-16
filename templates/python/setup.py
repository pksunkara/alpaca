import os
import sys

try:
    from setuptools import setup
except ImportError:
    from distutils.core import setup

setup(
	name='{{.Pkg.package}}',
	version='{{.Pkg.version}}',
	description='{{if .Pkg.official}}Official {{end}}{{.Pkg.name}} API library client for python',
	author='{{.Pkg.author.name}}',
	author_email='{{.Pkg.author.email}}',
	url='{{.Pkg.url}}',
	license='{{.Pkg.license}}',
	install_requires=[
		'requests'
	],
	packages=[
		'{{call .Fnc.underscore .Pkg.name}}'
	],
	classifiers=[
		'Development Status :: 5 - Production/Stable',
		'Intended Audience :: Developers',
		'License :: OSI Approved :: {{if .Pkg.python}}{{or .Pkg.python.license .Pkg.license}}{{end}}',
		'Operating System :: OS Independent',
		'Programming Language :: Python :: 2.7',
		'Topic :: Software Development :: Libraries :: Python Modules',
	],
	use_2to3 = True
)
