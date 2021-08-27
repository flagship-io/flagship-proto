#!/usr/bin/env bash
#### CONFIG ####
protoctag=0.4.0
ignore="^(bin|doc)/"
####

update_proto() {
    proto_dir_clean=${1%/}
    echo building $proto_dir_clean go files
    docker run --rm -v $(pwd):$(pwd) -w $(pwd) \
            znly/protoc:$protoctag \
            --go_out=$GOPATH/src/ \
            -I$GOPATH/src/ \
            -I. $proto_dir_clean/$proto_dir_clean.proto
}


if [ "$1" != "" ] && [ -d "$1" ]; then
    update_proto "$1"
else
    for proto_dir in $(ls -d */); do
        if [[ "$proto_dir" =~ $ignore ]]; then
            continue
        else
            update_proto "$proto_dir"
        fi
    done
fi
