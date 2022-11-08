#!/usr/bin/env sh

rm -r test 2> /dev/null
mkdir -p test/one test/two test/three
touch test/one/abc.txt test/one/bar.txt test/two/car.md test/two/door.md test/three/"hello world.md" test/three/"good bye.md"

rm -r auto 2> /dev/null
mkdir auto
echo "123 oo\n321 a" > auto/ap.txt
echo "789 oo\n987 e" > auto/pre.txt
echo "re a /-\nre oo __\nre c 1\nre b 2\nre oo __\ncase title kebab\ninsert '  ' 2 -c\ncase trim" > auto/auto.txt
