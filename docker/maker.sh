# !bin/sh

# # Needs to be executed from the root directory of the app

# Dependencies:
# npm:
#   -> axios
#   -> vue.js
# golang
#//// mysql / mariadb

# Need to be run from X/alpha-test

# FRONTEND
cd src/front-end

npm run build

# if the public folder already exists we remove it
if [ -d "../../public" ]; then
   rm -rf ../../public
fi
# and move the built vue.js folder into <app>/public
mv dist ../../public

# go to the bin folder of the app to set up the GOBIN
cd ../../bin
# use the app's bin folder
OLDGOBIN=$GOBIN
export GOBIN=`pwd`
# go back to the src folder
cd ../src
go install alpha.go
# clean footprints by setting back the GOBIN to what is was
export GOBIN=$OLDGOBIN

# Go back to the root directory of the app
cd ../bin
pwd
./alpha
