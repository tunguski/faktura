#!/bin/bash

#time docker run --rm -v "$PWD":/code -w /code/src golang /code/gobuild
docker run --rm -tiv `pwd`:/go mbrt/golang-vim-dev
