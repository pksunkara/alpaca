import os
import sys

try:
    from setuptools import setup
except ImportError:
    from distutils.core import setup

setup(
	name='{{.Pkg.Package}}',
	version='{{.Pkg.Version}}',
	description='{{if .Pkg.Official}}Official {{end}}{{.Pkg.Name}} API library client for python',
	author='{{.Pkg.Author.Name}}',
	author_email='{{.Pkg.Author.Email}}',
	url='{{.Pkg.Url}}',
	license='{{.Pkg.License}}',
	install_requires=[
		'requests >= 2.1.0'
	],
	packages=[
		'{{call .Fnc.underscore .Pkg.Name}}'
	],
	classifiers=[
		'Development Status :: 5 - Production/Stable',
		'Intended Audience :: Developers',{{if .Pkg.Python.License}}
		'License :: OSI Approved :: {{.Pkg.Python.License}}',{{end}}
		'Operating System :: OS Independent',
		'Programming Language :: Python :: 2.6',
		'Programming Language :: Python :: 2.7',
		'Programming Language :: Python :: 3.2',
		'Programming Language :: Python :: 3.3',
		'Topic :: Software Development :: Libraries :: Python Modules',
	]
)
