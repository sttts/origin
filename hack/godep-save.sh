#!/bin/bash

source "$(dirname "${BASH_SOURCE}")/lib/init.sh"

pin-godep() {
  pushd "${GOPATH}/src/github.com/tools/godep" > /dev/null
    git checkout "$1"
    "${GODEP}" go install
  popd > /dev/null
}

# build the godep tool
# Again go get stinks, hence || true
go get -u github.com/tools/godep 2>/dev/null || true
GODEP="${GOPATH}/bin/godep"

# Use to following if we ever need to pin godep to a specific version again
pin-godep 'v75'

# Some things we want in godeps aren't code dependencies, so ./...
# won't pick them up.
REQUIRED_BINS=(
  "github.com/elazarl/goproxy"
  "github.com/golang/mock/gomock"
  "github.com/containernetworking/cni/plugins/ipam/host-local"
  "github.com/containernetworking/cni/plugins/main/loopback"
  "k8s.io/kubernetes/cmd/libs/go2idl/go-to-protobuf/protoc-gen-gogo"
  "k8s.io/kubernetes/cmd/libs/go2idl/client-gen"
  "github.com/onsi/ginkgo/ginkgo"
  "github.com/jteeuwen/go-bindata/go-bindata"
  "./..."
)

TMPGOPATH=`mktemp -d`
trap "rm -rf $TMPGOPATH" EXIT
mkdir $TMPGOPATH/src

fork::without::vendor () {
  local PKG="$1"
  echo "Forking $PKG without vendor/"
  local DIR=$(dirname "$PKG")
  mkdir -p "$TMPGOPATH/src/$DIR"
  cp -a "$GOPATH/src/$PKG" "$TMPGOPATH/src/$DIR"
  pushd "$TMPGOPATH/src/$PKG" >/dev/null
    local OLDREV=$(git rev-parse HEAD)
    git rm -qrf vendor/
    git commit -q -m "Remove vendor/"
    local NEWREV=$(git rev-parse HEAD)
  popd >/dev/null
  echo "s/$NEWREV/$OLDREV/" >> "$TMPGOPATH/undo.sed"
}

fork::with::fake::packages () {
  local PKG="$1"
  shift
  echo "Forking $PKG with fake packages: $*"
  local DIR=$(dirname "$PKG")
  mkdir -p "$TMPGOPATH/src/$DIR"
  cp -a "$GOPATH/src/$PKG" "$TMPGOPATH/src/$DIR"
  pushd "$TMPGOPATH/src/$PKG" >/dev/null
    local OLDREV=$(git rev-parse HEAD)
    for FAKEPKG in "$@"; do
      mkdir -p "$FAKEPKG"
      echo "package $(basename $DIR)" > "$FAKEPKG/doc.go"
      git add "$FAKEPKG/doc.go"
    done
    git commit -a -q -m "Add fake package $*"
    local NEWREV=$(git rev-parse HEAD)
  popd >/dev/null
  echo "s/$NEWREV/$OLDREV/" >> "$TMPGOPATH/undo.sed"
}

undo::forks::in::godep::json () {
  echo "Replacing forked revisions with original revisions in Godeps.json"
  sed -f "$TMPGOPATH/undo.sed" Godeps/Godeps.json > "$TMPGOPATH/Godeps.json"
  mv "$TMPGOPATH/Godeps.json" Godeps/Godeps.json
}

fork::without::vendor github.com/docker/distribution
fork::without::vendor github.com/libopenstorage/openstorage
fork::with::fake::packages github.com/docker/docker \
  api/types \
  api/types/blkiodev \
  api/types/container \
  api/types/filters \
  api/types/mount \
  api/types/network \
  api/types/registry \
  api/types/strslice \
  api/types/swarm \
  api/types/versions

GOPATH=$TMPGOPATH:$GOPATH:$GOPATH/src/k8s.io/kubernetes/staging "${GODEP}" save -t "${REQUIRED_BINS[@]}"

undo::forks::in::godep::json
