
# Polish invoice creation system

## Goals

* Generate beautiful pdf invoices
* Respect all the rules defined by law
* Manage data as plain text
* Command line interface for operations
* Docker image containing all the prerequisites

## Command line interface

This application should conform to Unix philosophy.

## Docker image content

Final docker image should contain application itself
as well as latex toolchain for generating pdf files
from tex source files.

## Building application

Automated build script is defined:

```sh
./build.sh
```

## Tasks

* [ ] Enable golang modules
* [ ] Include cli, file processing, structured text parsing libraries
* [ ] Define initial structure for data
* [ ] Create initial template for invoice
* [ ] Understand correct role of git
* [ ] Creating docker image
    * [ ] Include Latex toolchain
    * [ ] Mount point for data

## External sources

* https://jakub.nadolny.info/faktura.tex
* https://ljvmiranda921.github.io/notebook/2018/04/23/postmortem-shift-to-docker/
