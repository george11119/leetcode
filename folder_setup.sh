#!/usr/bin/env bash

problem_name=$1
mkdir $problem_name
touch ${problem_name}/${problem_name}.go
touch ${problem_name}/${problem_name}_test.go
cd $problem_name
go mod init $problem_name
