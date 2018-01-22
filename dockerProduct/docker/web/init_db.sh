 #!/bin/bash
 #mysql --skip-grant-tables -u gotest -e "CREATE DATABASE ShopItems"
 mysql -u gotest < /tmp/dump.sql 