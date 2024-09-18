#!/bin/bash

java -jar ./antlr-4.13.1-complete.jar -Dlanguage=Go -encoding utf-8 -visitor -package parser -o ./alr *.g4
#antlr4 -Dlanguage=Go -encoding utf-8 -visitor -package parser -o ./alr *.g4