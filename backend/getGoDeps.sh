#!/bin/sh

# IF you are a windows user, this script won't work for you, har har har!

# DEPENDENCIES:
# go get github.com/hoisie/web
# go install log4go.googlecode.com/hg           ??
# go get github.com/go-sql-driver/mysql

function get_dependencies () {

  if [ "$GOPATH" = "" ]; then
    echo "you need to set your GOPATH"
    return
  fi

  echo "downloading go deps..."

  go get github.com/hoisie/web
  go get github.com/go-sql-driver/mysql
  go get code.google.com/p/log4go
  
  echo "...finished"
}

get_dependencies
