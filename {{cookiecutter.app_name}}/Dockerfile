# Note you cannot run golang binaries on Alpine directly
FROM            debian:buster-slim

MAINTAINER      {{cookiecutter.email}}

COPY            {{cookiecutter.app_name}} /{{cookiecutter.app_name}}

WORKDIR		/
ENV		GIN_MODE=release

EXPOSE          8080

ENTRYPOINT      [ "/{{cookiecutter.app_name}}" ]
