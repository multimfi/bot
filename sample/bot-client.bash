#!/usr/bin/env bash
readonly exe='bot-client'
readonly expire='5000'

function notify() {
	local notifyopts=(
		"--expire-time=$expire"
		"--urgency=$1"
		"$2" # summary
		"$3" # body
	)
	notify-send "${notifyopts[@]}"
}

"$exe"|while read -r state msg; do
	if [[ "$state" == 'firing' ]]; then
		notify 'critical' "$state" "$msg"
	elif [[ "$state" == 'resolved' ]]; then
		notify 'low' "$state" "$msg"
	fi
done
