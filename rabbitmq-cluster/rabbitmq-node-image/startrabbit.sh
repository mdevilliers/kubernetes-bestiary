#!/bin/bash

# if [ -z "$CLUSTERED" ]; then
# 	# if not clustered then start it normally as if it is a single server
# 	/usr/sbin/rabbitmq-server	
# else
 
chown rabbitmq:rabbitmq /var/lib/rabbitmq/.erlang.cookie
chmod 400 /var/lib/rabbitmq/.erlang.cookie

if [ -z "$RMQCONTROLLERSERVICE_SERVICE_HOST" ]; then
	# If clustered, but cluster with is not specified then again start normally, could be the first server in the
	# cluster
	/usr/sbin/rabbitmq-server
else
	/usr/sbin/rabbitmq-server -detached
	rabbitmqctl stop_app
	if [ -z "$RAM_NODE" ]; then
		rabbitmqctl join_cluster rabbit@$RMQCONTROLLERSERVICE_SERVICE_HOST
	else
		rabbitmqctl join_cluster --ram rabbit@$RMQCONTROLLERSERVICE_SERVICE_HOST
	fi
	rabbitmqctl start_app

	# Tail to keep the a foreground process active..
	tail -f /var/log/rabbitmq/rabbit\@$HOSTNAME.log
fi
# fi
