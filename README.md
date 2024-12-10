# mmdb2kuma

English version below

# Описание

Утилита предназначена для конвертации mmdb гео-базы в CSV файл, пригодный для загрузки в KUMA в качестве источника GeoIP информации.

# Использование

Файл __.mmdb__ в своем названии должен содержать `city` или `country` (регистр не важен) в зависимости от наполнения БД. По-умолчанию CSV генерируется на русском языке, для изменения языка необходимо передать аргумент `--lang` со значением соответствующего языка (например, `en`).

Для запуска можно воспользоваться командой:

```
go run mmdb2kuma.go GeoIP2-City.mmdb > ip.csv
```

Также можно скачать скомпилированный файл и запустить:

```cmd
.\mmdb2kuma.exe GeoIP2-City.mmdb > ip.csv
```

Или сделать билд самостоятельно:
```
go build mmdb2kuma.go
```

# What does it do?

Convert Maxmind mmdb database to CSV acceptable by KUMA.

# How?

go run mmdb2kuma GeoIP2City.mmdb > ips.csv

.\mmdb2kuma.exe GeoIP2City.mmdb > ips.csv

./mmdb2kuma GeoIP2City.mmdb > ips.csv

# Build

go build mmdb2kuma.go
