#!/bin/bash

# make xo models
host=95.163.202.160
db_name=homework
db_pass_word=vimi
./tools/bin/xo "mysql://root:${db_pass_word}@${host}/${db_name}?parseTime=true&sql_mode=ansi" -o ./db/model --template-path ./db/templates