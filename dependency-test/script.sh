#!bin/bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
echo "$SCRIPT_DIR"
go mod init hoang
if [ $? -ne 0 ]; then
  echo "Failed step 1"
  exit 1
fi
# go get if lack of package
go build
if [ $? -ne 0 ]; then
  echo "Failed step 3"
  echo "GO GET" 
  exit 1
fi
./hoang
  echo "Failed last step"
  exit 1
fi

exit 0 # return successful exit code.
