#!/bin/bash

URL=https://api.coindcx.com/exchange/ticker

curl -s $URL | jq '.[] | {ticker: .market, price: .last_price}'  | jq -r  '.ticker + " -> " + .price'