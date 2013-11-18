#!/bin/bash

FLAGCSV=flags.csv
USERSCSV=users.csv
OUTFILE="seed.sql"

function create_users_script () {
    local file=$OUTFILE

    OLDIFS=$IFS
    IFS=,
    while read PUBLICID NAME EMAIL
    do
        echo "INSERT INTO Users (publicId, name, email) VALUES ('$PUBLICID','$NAME','$EMAIL');" >> "$file"
    done < $USERSCSV
    IFS=$OLDIFS
}

function create_flags_script () {
    local file=$OUTFILE

    OLDIFS=$IFS
    IFS=,
    while read HOST HASH VALUE COMMENTS
    do
        echo "INSERT INTO Flags (host, tag, value, comment) VALUES ('$HOST', '$HASH', '$VALUE', '$COMMENTS');" >> "$file"
    done < $FLAGCSV
    IFS=$OLDIFS
}

> $OUTFILE
create_users_script
create_flags_script
