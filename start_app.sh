#!/bin/bash

# MySQLサーバが起動するまで待機
until mysqladmin ping -h mysql --silent; do
    echo 'waiting for mysql to be connectable ...'
    sleep 2
done

echo 'app is starting ...!'
exec go run app/main.go