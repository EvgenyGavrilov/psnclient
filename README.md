# Playstation store API client

[![GitHub Test Badge](https://github.com/EvgenyGavrilov/psnclient/workflows/Go/badge.svg)](https://github.com/EvgenyGavrilov/psnclient/actions)
[![codecov](https://codecov.io/gh/EvgenyGavrilov/psnclient/branch/main/graph/badge.svg?token=NGLW7QW30X)](https://codecov.io/gh/EvgenyGavrilov/psnclient)

Клиент для скачивания игр из магазина playstation.  
Включает в себя:
* получение списка игр
* получение полной информации об игре по идентификатору или заранее известному URL

Клиент создан на основе готового [psn-swagger](https://github.com/olegshulyakov/psn-swagger)

## Установка
```shell script
go get github.com/EvgenyGavrilov/psnclient
```

## Пример использования
**Получение списка игр**
```go
package main

import (
	"fmt"

	"github.com/EvgenyGavrilov/psnclient"
	"github.com/EvgenyGavrilov/psnclient/models"
)

func main() {
	cli, err := psnclient.New("ru", "ru")
	if err != nil {
		panic(err)
	}

	params := models.ListParams{
		Start: 0,
		Size:  2,
	}

	list, err := cli.ListGames(params)
	if err != nil {
		panic(err)
	}

	for _, game := range list.Links {
		fmt.Printf("Product ID: %s\n", game.ID)
	}
}
```

В результате получим следующие
```text
Product ID: EP1018-PPSA01696_00-BACK4BLOOD000000
Product ID: EP1018-PPSA01696_00-BACK4BLOODDELUXE
```

**Получение игры по ID или зарание известному URL**
```go
package main

import (
	"fmt"

	"github.com/EvgenyGavrilov/psnclient"
	"github.com/EvgenyGavrilov/psnclient/models"
)

func main() {
	cli, err := psnclient.New("ru", "ru")
	if err != nil {
		panic(err)
	}

	product, err := cli.ProductByID("EP1018-PPSA01696_00-BACK4BLOOD000000")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name: %s\n", product.Name)

	product, err = cli.ProductByURL("https://store.playstation.com/store/api/chihiro/00_09_000/container/RU/ru/999/EP1018-PPSA01696_00-BACK4BLOOD000000/1617052056")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name: %s\n", product.Name)
}
```
В результате получим следующие
```text
Product name by ID: Back 4 Blood: Стандартное издание PS4 and PS5
Product name by URL: Back 4 Blood: Стандартное издание PS4 and PS5
```
