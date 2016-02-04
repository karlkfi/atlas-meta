# Atlas-Meta

Tool for updating Atlas (vagrant-cloud) box repository JSON files.

Atlas-Meta allows programmatically adding new vagrant box versions to an Atlas (formerly vagrant-cloud) "Box Catalog Metadata" file. This metadata file can then be specified as a `box_url` in a Vagrantfile, which allows vagrant to automatically download the latest version of a box.

Atlas-Meta is built into a docker image to make it easy to use in continuous integration without installation: https://hub.docker.com/r/karlkfi/atlas-meta/

Vagrant Box Metadata Format: https://www.vagrantup.com/docs/boxes/format.html


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


## Build Docker Image

```
./build.sh
```


## License

Copyright 2016 Karl Isenberg

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
