# Docker: Golang web + MySQL


# Run containers
docker-compose build
docker-compose up -d


mysql -uroot --port="3307" -h127.0.0.1 -p"gotest" -e "CREATE DATABASE ShopItems"
mysql -uroot --port="3307" -h127.0.0.1 -pgotest ShopItems < docker/db/dump.sql


# Confirm it all went up correctly. Exit with Ctrl+C
docker-compose logs
```

### Starting  app
 
```
docker exec -it dockerproduct_web_1 bash
 
