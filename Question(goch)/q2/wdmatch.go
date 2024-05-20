package main

import (
"os"
)

func wdMatch(a, b string) string {
if len(os.Args) != 2 {
return
}
a = os.Args[1]
b = os.Args[2]

for _, char1 := range a{
for  _, char2 := range b {

