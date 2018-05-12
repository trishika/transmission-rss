all: transmission-rss

prepare:
	go get

transmission-rss: $(wildcard *.go)
	go build -o $@

clean:
	rm -f transmission-rss
