#!/usr/bin/env bash

p1() {
	rg 'mul\((\d+),(\d+)\)' "$1" -oN | sed -E 's/mul\(([0-9]+),([0-9]+)\)/\1*\2/' | bc | xargs echo | sed 's/ /+/g' | bc
}

p2() {
	local state=1
	local str="0"
	for i in $(rg "(mul\((\d+),(\d+)\)|do\(\)|don\'t\(\))" "$1" -oN); do
		case "$i" in
		"don't()")
			state=0
			;;
		"do()")
			state=1
			;;
		*)
			if [ $state == 1 ]; then
				str="${str}+$(echo "$i" | sed -E 's/mul\(([0-9]+),([0-9]+)\)/\(\1*\2\)/')"
			fi
			;;
		esac
	done
	echo "$str" | bc
}

p1 "$@"
p2 "$@"
