#!/usr/bin/env -S sed -n -f
s/^\(\..*\)\/.*:\/\/go:generate go tool stringer -type \([a-zA-Z][^ ]\+\) .*$/\1 \2/p
