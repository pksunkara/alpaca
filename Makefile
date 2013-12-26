all:
	alpaca tests

deps: all
	go get github.com/wsxiaoys/terminal
	go get github.com/codegangsta/martini

ruby:
	gem install faraday -v 0.8.8 --no-ri --no-rdoc
	gem install json -v 1.7.7 --no-ri --no-rdoc

node:
	npm install request 2.x.x
	npm install async 0.2.x

php:
	composer require guzzle/guzzle 3.7.*

python:
	pip install "requests>=2.1.0"

clean:
	rm -rf tests/node tests/python tests/php tests/ruby
	rm -rf node_modules vendor composer.json composer.lock
