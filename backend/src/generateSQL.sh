#!/bin/bash

FLAGCSV=flags.csv
USERSCSV=users.csv
OUTFILE="seed.sql"

function create_users_script () {

    echo "-- Users" >> "$OUTFILE"
    local count=0

    OLDIFS=$IFS
    IFS=,
    while read PUBLICID NAME EMAIL
    do
        if [ $count -ne 0 ]; then
            echo "INSERT INTO Users (publicId, name, email) VALUES ('$PUBLICID','$NAME','$EMAIL');" >> "$OUTFILE"
        fi
        let "count += 1"
    done < $USERSCSV
    IFS=$OLDIFS
}

function create_flags_script () {

    echo "-- Flags" >> "$OUTFILE"
    local count=0

    OLDIFS=$IFS
    IFS=,
    while read HOST HASH VALUE COMMENTS
    do
        if [ $count -ne 0 ]; then
            echo "INSERT INTO Flags (host, tag, value, comment) VALUES ('$HOST', '$HASH', '$VALUE', '$COMMENTS');" >> "$OUTFILE"
        fi
        let "count += 1"
    done < $FLAGCSV
    IFS=$OLDIFS
}

> $OUTFILE
create_users_script
echo "" >> "$OUTFILE"
create_flags_script
