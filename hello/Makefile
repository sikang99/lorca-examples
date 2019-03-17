#
# Makefile for lorca
#
PROG=hello
usage:
	@echo "make [build|run|clean|dev]"

edit e:
	vi main.go

build b:
	go generate
	go build -o $(PROG)

run r:
	./$(PROG)

clean c:
	rm -f $(PROG) assets.go

# ---------------------------------------------------
BROWSER=chromium-browser 	# chrome.exe
dev d:
	$(BROWSER) --remote-debugging-port=9222 &
	xdg-open http://localhost:9222/json
