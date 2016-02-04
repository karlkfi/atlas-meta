# Atlas-Meta

Tool for updating Atlas (vagrant-cloud) box repository JSON files.

## Install

```
go get github.com/golang/glog
go install github.com/karlkfi/atlas-meta
```

## Build Docker Image

```
./build.sh
```


## Example

```
atlas-meta \
  --repo ~/workspace/dcos-vagrant/build/metadata.json \
  --version 0.3.0 \
  --status active \
  --desc 'Automated CI Build' \
  --provider virtualbox \
  --box https://s3-us-west-1.amazonaws.com/dcos-vagrant/dcos-centos-virtualbox-0.3.0.box \
  --checksum-type sha1 \
  --checksum fake-checksum \
  add
```

```
docker run --rm -it \
  -v "$HOME/workspace/dcos-vagrant:/dcos-vagrant" \
  karlkfi/atlas-meta \
  --repo /dcos-vagrant/build/metadata.json \
  --version 0.3.0 \
  --status active \
  --desc 'Automated CI Build' \
  --provider virtualbox \
  --box file:///~/workspace/dcos-vagrant/build/dcos-centos-virtualbox-0.3.0-alpha.box \
  --checksum-type sha1 \
  --checksum fake-checksum \
  add
```