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

makeSymbolsFile()
processSymbols()