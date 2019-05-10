#!/bin/bash


###
echo ""
echo "Pushing changed to repository..."
echo ""
git add .; git commit -m "automatic from deployment script"; git push

# go build -o greenguard &&
# CGO_ENABLED=1 GOARCH="arm64" GOOS="linux" go build -o greenguard_arm &&

# scp -r $GOPATH/src/bitbucket.org/vacovsky/greenguard joe@192.168.111.127:/opt/ > /dev/null


rsync -av -e ssh --exclude='.git' $GOPATH/src/bitbucket.org/vacovsky/greenguard joe@192.168.111.116:/opt/

# scp -r $GOPATH/src/bitbucket.org/vacovsky/greenguard_arm joe@192.168.111.127:/opt/greenguard/greenguard > /dev/null
