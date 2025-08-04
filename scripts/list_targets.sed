#!/usr/bin/env -S sed -n -f

s/^## \(.*\): \(.*\)/\1|\2/p
