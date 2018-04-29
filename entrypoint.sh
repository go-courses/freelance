#!/bin/bash
echo "Waiting for mysql"
until mysql -hfreelance-db -P3306 -uroot -proot &> /dev/null
do
  printf "."
  sleep 1
done

echo -e "\nmysql ready"

/freelance