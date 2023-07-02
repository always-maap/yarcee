#!/sbin/openrc-run

name=$RC_SVCNAME
description="frontline api"
supervisor="supervise-daemon"
command="/usr/local/bin/frontline"
pidfile="/run/fire.pid"
command_user="frontlineApi:frontlineApi"

depend() {
	after net
}