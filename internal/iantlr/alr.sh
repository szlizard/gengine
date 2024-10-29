#!/bin/bash

VERSION=4.13.2
java -jar ./antlr-${VERSION}-complete.jar -Dlanguage=Go -encoding utf-8 -visitor -package parser -o ./alr *.g4
#antlr4 -Dlanguage=Go -encoding utf-8 -visitor -package parser -o ./alr *.g4