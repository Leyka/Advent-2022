#!/usr/bin/env bash

last_day=$(ls | grep ^day | grep -Eo '[0-9]*' | tail -1)
last_day_increment=$((last_day + 1))

new_day_padded=$(printf "%02d\n" $last_day_increment)
new_day_dirname="day$new_day_padded"

mkdir "$new_day_dirname"
cp -a template/ "$new_day_dirname"
