#!/bin/bash

CUR_DIR=$(pwd)
MYSQL_USER='root'
MYSQL_PSW='123456'

function check_app {
  appName=$1
  word=""
  if [[ "$appName" = "linktime-mysql" ]]; then
    word="socket: '/var/run/mysqld/mysqld.sock'  port: 3306"
  fi
  echo "Checking if $appName is ready ..."
  line=`docker logs default-$appName 2>&1 | grep "${word}" | wc -l`
  while [ $line != "1" ]
  do
    sleep 2
    line=`docker logs default-$appName 2>&1 | grep "${word}" | wc -l`
  done
  echo "$appName is ready!"
}

function start_required_service {
  appName=$1

  if [[ $appName =~ "linktime-mysql" ]]; then
    docker rm -f default-$appName
    docker run --name default-$appName -p 3306:3306 -e MYSQL_ROOT_PASSWORD=$MYSQL_PSW -d mysql:5.7.18
    check_app $appName
  fi
}

cd $CUR_DIR

mkdir -p upload

# start_required_service "linktime-mysql"

running_status=`docker ps --filter "name=default-linktime-mysql" | wc -l`
if [ $running_status -ne 0 ] || [ $restartDocker -ne 0 ]; then
  echo "creating database..."
  docker cp init.sql default-linktime-mysql:/.
  docker exec default-linktime-mysql /bin/sh -c "mysql -u $MYSQL_USER -p$MYSQL_PSW < /init.sql"
  echo "create database successfully"
fi

./backend
