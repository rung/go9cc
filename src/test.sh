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

try 0 'main(){return 0;}'
try 42 'main(){return 42;}'
try 21 'main(){return 5+20-4;}'
try 41 "main(){return  12 + 34 - 5 ;}"
try 47 "main(){return 5+6*7;}"
try 15 "main(){return 5*(9-6) ;}"
try 4 "main(){return (3+5)/2;}"
try 4 "main(){return (3+5)/2; }"
try 4 "main(){return a=(3+5)/2;a;}"
try 10 "main(){return z=5*2;z;}"
try 10 "main(){test=5*2;return test;}"
try 1 "main(){test= 5==5;return test;}"
try 0 "main(){test= 5==6;return test;}"
try 0 "main(){test= 5!=5;return test;}"
try 1 "main(){test= 5!=6;return test;}"
try 0 "main(){foo(); return 0;}"
try 0 "main(){a=1; bar(a,2,3); return 0;}"
try 0 "main(){b=1; barr(1,b,3,4,5,6,7,8,9); return 0;}"
try 3 "main(){b=3;return b; c=5; return c;}"
try 1 "sub(){return 1;} main(){return sub();}"
try 15 "sub(){return 15;} main(){return sub();}"
try 20 "sub(){return 15;} main(){return sub()+5;}"
try 0 "main(){a=1+3;}"

echo OK
