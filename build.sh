#!/bin/bash

time docker run --rm -v "$PWD":/code -w /code/src golang /code/gobuild
