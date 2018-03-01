#!/bin/sh

### BEGIN INIT INFO
# Provides:          slackbot
# Required-Start:    $local_fs $local_fs
# Required-Stop:     $local_fs $local_fs
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: golang slackbot
### END INIT INFO

DAEMON=/usr/local/bin/slackbot
NAME=slackbot
DESC=slackbot

PIDFILE=/var/run/$NAME.pid

. /lib/lsb/init-functions

case "$1" in
  start)
    test -d ${SLACKBOT_RUN_DIR:-/var/run/slackbot} || mkdir -p ${SLACKBOT_RUN_DIR:-/var/run/slackbot}
    log_daemon_msg "Starting $DESC: $NAME"
    start-stop-daemon --start --background  --pidfile $PIDFILE --make-pidfile --exec $DAEMON
    log_end_msg $?
    ;;
  stop)
    log_daemon_msg "Stopping golang slackbot: $NAME"
    start-stop-daemon --stop --pidfile $PIDFILE --retry 10 #--exec $DAEMON
    log_end_msg $?
    ;;
  restart)
    test -d ${SLACKBOT_RUN_DIR:-/var/run/slackbot} || mkdir -p ${SLACKBOT_RUN_DIR:-/var/run/slackbot}
    log_daemon_msg "Restarting $DESC: "
    start-stop-daemon --stop --quiet --pidfile $PIDFILE --exec $DAEMON
    sleep 1
    start-stop-daemon --start --background --pidfile $PIDFILE --exec $DAEMON
    log_end_msg $?
    ;;
  reload)
    log_daemon_msg "Reloading $DESC: "
    start-stop-daemon --stop --quiet --pidfile $PIDFILE --exec $DAEMON
    sleep 1
    start-stop-daemon --start --background --pidfile $PIDFILE --exec $DAEMON
    log_end_msg $?
    ;;
  status)
    status_of_proc "$DAEMON" "$NAME" && exit 0 || exit $?
    ;;
  *)
  echo "Usage: $NAME {start|stop|restart|reload|status}"
  exit 1
  ;;
esac

exit 0
