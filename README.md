# go-rest-api
Simple REST API written as an introduction to programming in Go!

# scrape
The contents of the scrape folder use mainly the io/ioutil and net/http modules
to access a url and get the text of its title HTML tag.

The program will panic (abruptly exit) if a URL not beginning with "http" or "https" is entered

# api 
The api folder uses the config.json file, which contains a host and port for ourAPI to use.

It has three endpoints (pages the API can go to and displau some information): the homepage "/", "/helloworld", and "/status". /status will display the 200 OK code, / will display HTML for the homepage, and /helloworld will display the "Hello world" greeting to the user
