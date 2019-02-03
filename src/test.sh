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

try 0 '0;'
try 42 '42;'
try 21 '5+20-4;'
try 41 " 12 + 34 - 5 ;"
try 47 "5+6*7;"
try 15 "5*(9-6) ;"
try 4 "(3+5)/2;"
try 4 "(3+5)/2; "
try 4 "a=(3+5)/2;a;"
try 10 "z=5*2;z;"
try 10 "test=5*2;test;"
try 1 "test= 5==5;test;"
try 0 "test= 5==6;test;"
try 0 "test= 5!=5;test;"
try 1 "test= 5!=6;test;"
try 0 "foo();"
try 0 "a=1; bar(a,2,3);"
try 0 "b=1; barr(1,b,3,4,5,6,7,8,9);"

echo OK
