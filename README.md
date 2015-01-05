# warmup

Simple HTTP cache warming for multiple URLs with headers support.

## Installation

```
$ go get -u github.com/frozzare/warmup
```

## The URLs JSON file

A simple JSON file with the url as key in the root object and a object with headers, for example Cookie header. If you don't need headers just leave a empty object.

```javascript
{
	"http://example.com": {
		"Cookie": "cookie body"
	}
}
```

## Usage of warmup

```
  -delay=29000: Delay (in ms) between items in the request list
  -filename=urls.json: List of URLs
  -timeFile=true: If you want warmup.txt beeing created with last run time
```
# License

MIT © [Fredrik Forsmo](https://github.com/frozzare)
