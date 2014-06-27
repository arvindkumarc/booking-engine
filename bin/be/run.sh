#!/bin/sh
SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)
"$SCRIPTPATH/booking-engine" -importPath github.com/arvindkumarc/booking-engine -srcPath "$SCRIPTPATH/src" -runMode prod
