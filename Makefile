# yt-channel-watcher
# See LICENSE for copyright and license details.
.POSIX:

PREFIX ?= /usr
GO ?= go
GOFLAGS ?= -buildvcs=false
RM ?= rm -f

all: yt-channel-watcher

yt-channel-watcher:
	$(GO) build $(GOFLAGS) .

install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f yt-channel-watcher $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/yt-channel-watcher

uninstall:
	$(RM) $(DESTDIR)$(PREFIX)/bin/yt-channel-watcher

clean:
	$(RM) yt-channel-watcher

.PHONY: all yt-channel-watcher install uninstall clean
