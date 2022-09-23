#!/sbin/openrc-run

name=$RC_SVCNAME
description="firerunner agent"
supervisor="supervise-daemon"
command="/usr/local/bin/fire"
pidfile="/run/fire.pid"
command_user="firerunner:firerunner"

depend() {
	after net
}