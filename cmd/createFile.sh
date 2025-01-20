#!/bin/bash

file=".env"

if ! [ -f $file ]; then
    touch $file
    echo "DEV_API_KEY=test" >> $file;
    echo "DEBUG=true" >> $file;

    echo ""
    echo "The data file was created"
    echo ""
fi
