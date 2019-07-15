# making a json to load in server and append actual names
# to all requests
# also used the verify symbols before get requests

import requests
import json

def makeSymbolsFile():
	url = "https://api.iextrading.com/1.0/ref-data/symbols"

	response = requests.get(url)

	print(response.json())

	with open("allSymbols.json", "w") as jsonOutfile:
		json.dump(response.json(), jsonOutfile)

def processSymbols():

	theHash = {}
	
	with open("symbols.json") as jsonInfile:

		data = json.load(jsonInfile)
		for symbol in data:
			theHash[symbol["symbol"]] = symbol["name"]

	with open("symbolHash.json", "w") as jsonOutfile:
		json.dump(theHash, jsonOutfile)
	
	print(theHash)

def getOpeningPrices():

	openingPriceCache = {}

	with open("symbolHash.json") as jsonInFile:

		allSymbols = json.load(jsonInFile)
		openingPriceCache = {}
		count = 0
		whichSubfile = 1
		for symbol in allSymbols:
			url = F'https://cloud.iexapis.com/beta/stock/{symbol}/quote/open?token=sk_5c1a1ec78f534b179da588b787245fe6'
			response = requests.get(url)

			if response is not None: 
				# print(symbol, response.json())
				openingPriceCache[symbol] = response.json()
		
			count += 1

			if count == 50:

				with open(F'openingPriceCache-{whichSubfile}.json', "w") as jsonOutfile:
					json.dump(openingPriceCache, jsonOutfile)

				openingPriceCache = {}
				count = 0
				whichSubfile += 1

# makeSymbolsFile()
# processSymbols()
getOpeningPrices()

