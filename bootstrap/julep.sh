#!/bin/bash

confdir=${1:-$HOME/julep}

[[ ! -d $confdir ]] && mkdir $confdir 
cd $confdir
echo `pwd`
exit 0
function hasCommand() {
    progName=$1
    [[ 1 -eq $(which $progName | head -n 1 | wc -l | awk '{print $1}') ]] && \
        echo true || \
        echo false 
}
function verifyCommand() {
    progName=$1
    if [[ $(hasCommand $progName) == false ]]; then
        echo "$progName is not found in your path. Bailing." 
        exit 1
    fi
}

verifyCommand go

verifyCommand gpg2

verifyCommand git

if [[ ! $(env | grep GOPATH) ]]; then
	echo '$GOPATH not set. Bailing....'
	exit 1
fi

if [[ ! $(echo $PATH | grep $GOPATH) ]]; then
	echo '$GOPATH is not in your $PATH. Bailing....'
	exit 1
fi

[[ ! -f .secring.gpg && ! -f .pubring.gpg ]] && \
	gpg2 --batch --armor --gen-key app.batch

[[ $(hasCommand etcd) != true ]] && \
	echo Installing etcd && \
	git clone https://github.com/coreos/etcd.git && \
	pushd etcd && \
	bash build && \
	cp bin/etcd bin/etcdctl $GOPATH/bin && \
	popd && \
	rm -rf etcd

[[ $(hasCommand crypt) == false ]] && \
	echo Installing crypt && \
	go get github.com/xordataexchange/crypt/bin/crypt

[[ ! $(screen -ls | grep '.julep-etcd') ]] && \
	screen -dmS julep-etcd etcd 

while [[ $(etcdctl ls 2>&1 | grep Error | wc -l) -gt 0 ]]; do
	sleep 1
done

if [[ $(etcdctl get /julep/config.json 2>&1 | grep 'Key not found' | wc -l) -eq 1 ]]; then
	crypt set -keyring .pubring.gpg /julep/config.json config.json
fi
