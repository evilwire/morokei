#! /bin/sh

# build the binary with a certain version
build() {
    echo "--> Building morokei (version ${VERSION})"
    if [[ ! -d ${VERSION} ]]; then
        mkdir /go/out/${VERSION}
    fi

    go install && cp /go/bin/morokei /go/out/${VERSION}/morokei
}

build $1