SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go' | grep -v '/main_')
ALL = parse simulate download
all: ${ALL}

clean:
	rm -f ${ALL}

#BUILD_TIME=`date +%FT%T%z`
LDFLAGS=

.DEFAULT_GOAL: $(BINARY)

download: $(SOURCES) main_download.go
	go build ${LDFLAGS} -o $@ ${SOURCES} main_$@.go

parse: $(SOURCES) main_parse.go
	go build ${LDFLAGS} -o $@ ${SOURCES} main_$@.go

simulate: $(SOURCES) main_simulate.go
	go build ${LDFLAGS} -o $@ ${SOURCES} main_$@.go

install-packages:
	go get github.com/PuerkitoBio/goquery
	go get github.com/go-sql-driver/mysql
	go get github.com/spf13/viper
