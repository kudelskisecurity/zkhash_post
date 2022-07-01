#!/bin/sh

addchain search '20974350070050476191779096203274386335076221000211055129041463479975432473805' > inv.acc
addchain gen -tmpl template_gnark inv.acc | gofmt > fun.gnark
