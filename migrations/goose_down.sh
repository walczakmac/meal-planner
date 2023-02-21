#!/bin/zsh
goose postgres "host=0.0.0.0 port=5432 user=fav_food password=fav_food dbname=fav_food sslmode=disable" down