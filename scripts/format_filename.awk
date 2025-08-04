#!/usr/bin/env -S awk -f

{ print sprintf("%s_string.go", tolower($0)) }
