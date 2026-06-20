#!/usr/bin/env bash

use_tree=false
problem_name=""

for arg in "$@"; do
	case $arg in
		--tree)
			use_tree=true
			;;
		*)
			problem_name=$arg
			;;
	esac
done

script_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

mkdir $problem_name
touch ${problem_name}/${problem_name}.go
cd $problem_name
go mod init $problem_name
echo package ${problem_name} >> ${problem_name}.go

if [ "$use_tree" = true ]; then
	sed "s/^package .*/package ${problem_name}/" "${script_dir}/test_templates/binary_tree_test_template.go" > ${problem_name}_test.go
else
	touch ${problem_name}_test.go
	echo package ${problem_name} >> ${problem_name}_test.go
fi
