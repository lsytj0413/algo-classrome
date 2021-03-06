#!/bin/bash

usage() {
    echo "$PROGNAME: usage: $PROGNAME [-s src] [-d dest] [--dir jsdir]"
    return
}

PROGNAME=$(basename $0)
BASEPATH=$(cd `dirname $0`; pwd)
src=
dest=
filename=
dir=

while [[ -n $1 ]]; do
    case $1 in
    -s)             shift
                    src=$1
                    ;;
    -d)             shift
                    dest=$1
                    ;;
    --dir)          shift
                    dir=$1
                    ;;
    -h | --help)    usage
                    exit
                    ;;
    *)              usage >&2
                    exit 1
                    ;;
    esac
    shift
done

if [[ (-z $src) || (-d $dest) ]]; then
    usage >&2
    echo "-s and -d must be provided"
    exit 1
fi

if [[ ! (( "$src" =~ ^[0-9]+$ ) && ( "$dest" =~ ^[0-9]+$ )) ]]; then
    echo "-s and -d must be integer"
    exit 1
fi

if [[ -z $dir ]]; then
    dir=$(dirname ${BASEPATH})/js
    echo "--dir not provided, use default result ${dir}"
fi

echo "cpp from ${dir}/p${src} to ${dir}/p${dest} "

cp -r ${dir}/p${src} ${dir}/p${dest}
mv ${dir}/p${dest}/p${src}.js ${dir}/p${dest}/p${dest}.js
mv ${dir}/p${dest}/p${src}.test.js ${dir}/p${dest}/p${dest}.test.js

sed -i'' "s/${src}/${dest}/g" ${dir}/p${dest}/p${dest}.js
sed -i'' "s/${src}/${dest}/g" ${dir}/p${dest}/p${dest}.test.js

echo "Done"
exit