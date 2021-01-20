#!/bin/bash

migrations_path="migrations"
today=`date +%Y%m%d%H%M%S`

touch "${migrations_path}/${today}_$1.down.sql"
touch "${migrations_path}/${today}_$1.up.sql"