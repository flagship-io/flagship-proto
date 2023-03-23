#!/usr/bin/env bash
#### CONFIG ####
ignore="^(bin|doc)/"
gitrepo="github.com\/flagship-io\/flagship-proto"
####

update_proto() {
    protoname=${1%/}
    rootpath=`echo $(pwd) | sed -e "s/$gitrepo//g"`
    echo building $protoname go file
    docker run --rm -v $(pwd):$(pwd) -w $(pwd) namely/protoc-all -f $protoname/$protoname.proto -l go -o $rootpath
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
