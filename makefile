
GIT_BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
GIT_HASH = $(shell git rev-parse --short HEAD)
GIT_TIME = $(shell date +%s)

all:	clean get test

clean:
	rm -rf src/google.golang.org src/golang.org

get:
	git clone git@github.com:DramaFever/golang.org-x-net --branch master src/golang.org/x/net && cd src/golang.org/x/net && git checkout 04b9de9b512f58addf28c9853d50ebef61c3953e
	git clone git@github.com:DramaFever/appengine.git src/google.golang.org/appengine && cd src/google.golang.org/appengine && git checkout e234e71924d4aa52444bc76f2f831f13fa1eca60

test:
	echo ${PWD}
	GOPATH=${PWD} goapp test starter
