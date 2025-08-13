#!/usr/bin/env -S awk -f

{ 
	split($0, a, " ");
	printf("%s/%s_string.go\n", a[1], tolower(a[2]))
}
