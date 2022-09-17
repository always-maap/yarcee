#!/sbin/openrc-run

name=$RC_SVCNAME
description="CodeBench agent"
supervisor="supervise-daemon"
command="/usr/local/bin/fire"
pidfile="/run/fire.pid"
command_user="codebench:codebench"

depend() {
	after net
}