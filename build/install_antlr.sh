#!/bin/bash

ANTLR_VERSION=antlr-4.10.1-complete

wget https://www.antlr.org/download/$ANTLR_VERSION.jar && \
sudo mv $ANTLR_VERSION.jar /usr/bin/$ANTLR_VERSION.jar
