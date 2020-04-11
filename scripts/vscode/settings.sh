#!/bin/bash

if [ $# -ne 2 ]
then
    echo "Pass path to '.vscode' directory as the first argument, \
'path-to-project' as the second argument."
    exit 1
fi

PATH_TO_VSCODE_DIR=$1
VSCODE_SETTINGS_FILEPATH=settings.json
ABSL_PATH_TO_PROJECT=$(readlink -f $2)
ABSL_PATH_TO_VENDOR_DIR=$ABSL_PATH_TO_PROJECT/vendor

cd $PATH_TO_VSCODE_DIR 

[[ -f $VSCODE_SETTINGS_FILEPATH ]] && touch $VSCODE_SETTINGS_FILEPATH

cat > $(readlink -f $VSCODE_SETTINGS_FILEPATH) << 'EOF' 
{
    "go.gopath": "__ABSL_PATH_TO_PROJECT__:__ABSL_PATH_TO_VENDOR_DIR__",
    "terminal.integrated.env.linux": {
        "GOPATH": "${go.gopath}",
        "GOBIN": "${go.gopath}/bin"
    }
}
EOF

sed -i -e "s@__ABSL_PATH_TO_PROJECT__@$ABSL_PATH_TO_PROJECT@g" $VSCODE_SETTINGS_FILEPATH
sed -i -e "s@__ABSL_PATH_TO_VENDOR_DIR__@$ABSL_PATH_TO_VENDOR_DIR@g" $VSCODE_SETTINGS_FILEPATH

cd - > /dev/null