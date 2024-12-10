# mmdb2kuma

# What does it do?

Convert Maxmind mmdb database to CSV.

# Why?

Many applications support CSV but not mmdb.  For example it's easy to import CSV to SQL databases.

# How?

go run mmdb2kuma GeoIP2City.mmdb > ips.csv

.\mmdb2kuma.exe GeoIP2City.mmdb > ips.csv

./mmdb2kuma GeoIP2City.mmdb > ips.csv

# Build
go build mmdb2kuma.go

