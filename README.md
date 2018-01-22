# productTask

[Запуск через Docker - инструкция](dockerProduct/README.md)

## Точки входа

1. /v1/PostProduct - отправка данных продуктов магазина
2. /v1/GetProduct?id= - получение данных конкретного продукта по id, например /v1/GetProduct?id=1


## Задача

Можно использовать любые библиотеки.
Результат необходимо залить в любой репозиторий.

- Написать сервис на Golang, который принимает массив URL-ов товаров https://www.amazon.co.uk,
для данных URL он должен загрузить наименование, цена, фото товара(только URL, грузить изображение не надо), признак наличия, пример:

``` json
[
  {
    "url": "https://www.amazon.co.uk/gp/product/1509836071",
    "meta": {
        "title": "The Fat-Loss Plan: 100 Quick and Easy Recipes with Workouts",
        "price": "8.49",
        "image": "https://images-na.ssl-images-amazon.com/images/I/51kB2nKZ47L._SX382_BO1,204,203,200_.jpg",
    }
  },
  // ...
]
```
- Сервис необходимо завернуть в Docker
- [Не обязательно] Плюсом будет реализация асинхронной загрузки, когда по некоторому requestID мы можем получить результат
