#!/bin/bash

# make xo models
host=127.0.0.1
db_name=homework
db_pass_word=dc
./tools/bin/xo "mysql://root:${db_pass_word}@${host}/${db_name}?parseTime=true&sql_mode=ansi" -o ./db/model --template-path ./db/templates