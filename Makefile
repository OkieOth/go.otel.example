.PHONY: test build version

VERSION = $(shell cat version.txt)

version:
	cat version.txt

start-dev-env:
	bash -c "docker/bin/compose_env.sh start"

stop-dev-env:
	bash -c "docker/bin/compose_env.sh destroy"
