#! /bin/sh

# build the binary with a certain version
build() {
    echo "--> Building morokei (version $1)"
    if [[ ! -d $1 ]]; then
        mkdir /go/out/$1
    fi

    go install && cp /go/bin/morokei /go/out/$1/morokei
}

build $1