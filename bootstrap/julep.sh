#!/bin/bash

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

for i in $(echo $GOPATH | sed 's/:/ /g'); do
	if [[ ! $(echo $PATH | grep $i) ]]; then
		echo "$i is not in your \$GOPATH"
		exit 1
	fi
done

[[ ! -f .secring.gpg && ! -f .pubring.gpg ]] && \
	gpg2 --batch --armor --gen-key app.batch

[[ $(hasCommand etcd) != true ]] && \
	echo Installing etcd && \
	git clone https://github.com/coreos/etcd.git && \
	pushd etcd && \
	bash build && \
	cp bin/etcd bin/etcdctl $(echo $GOPATH | sed 's/:/ /g' | awk '{print $1}')/bin && \
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


if [[ ! -e julepEnv.sh ]]; then
	echo export JULEP_ETCD_URL="http://localhost:4001/" >> julepEnv.sh
	echo export JULEP_PRIVATE_KEY=`pwd`.secring.gpg >> julepEnv.sh
	echo "Run \"source julepEnv.sh\" to get started!"
fi
