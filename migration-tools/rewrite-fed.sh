#!/bin/bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e

sr() (
  perl -p -i -e "s|$1|$2|g" `grep -ril $1 *`
)

pushd federation > /dev/null
  sr 'kubernetes/federation' federation
  sr 'kubernetes\.federation' federation
popd
git mv federation/Makefile Makefile.federation
git mv federation/BUILD BUILD
git mv federation/test .
git mv federation/cluster test/
git mv federation/* .
rmdir federation

git add .
