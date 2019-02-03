#!/bin/bash

try() {
  expected="$1"
  input="$2"

  ./9cc "$input" > tmp.s
  gcc -o tmp tmp.s testprint.o
  ./tmp
  actual="$?"

  if [ "$actual" = "$expected" ]; then
    echo "$input => $expected"
  else
    echo "$expected expected, but got $actual"
    exit 1
  fi
}

try 0 'main(){0;}'
try 42 'main(){42;}'
try 21 'main(){5+20-4;}'
try 41 "main(){ 12 + 34 - 5 ;}"
try 47 "main(){5+6*7;}"
try 15 "main(){5*(9-6) ;}"
try 4 "main(){(3+5)/2;}"
try 4 "main(){(3+5)/2; }"
try 4 "main(){a=(3+5)/2;a;}"
try 10 "main(){z=5*2;z;}"
try 10 "main(){test=5*2;test;}"
try 1 "main(){test= 5==5;test;}"
try 0 "main(){test= 5==6;test;}"
try 0 "main(){test= 5!=5;test;}"
try 1 "main(){test= 5!=6;test;}"
try 0 "main(){foo();}"
try 0 "main(){a=1; bar(a,2,3);}"
try 0 "main(){b=1; barr(1,b,3,4,5,6,7,8,9);}"

echo OK
