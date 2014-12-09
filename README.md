# go-warmup

Simple HTTP cache warming for multiple URLs with headers support.

## Installation

```
$ go get -u github.com/frozzare/go-warmup
```

## The URLs JSON file

A simple JSON file with the url as key in the root object and a object with headers, for example Cookie header. If you headers needed just leave a empty object.

```javascript
{
	"http://example.com": {
		"Cookie": "cookie body"
	}
}
```

## Usage of warmup

```
  -delay=29000: Delay (in ms) between request lists
  -filename=urls.json: List of URLs
  -timeFile=true: If you want warmup.txt beeing created with last run time
```
# License

Copyright (c) 2014 Fredrik Forsmo
Licensed under the MIT license.
