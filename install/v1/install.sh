#!/bin/bash
##### @jeffotoni
DIR=$GOPATH/bin
EXEC=gobootcamp
ZSHRC=$HOME/.zshrc
BASHRC=$HOME/.bashrc

rm -f $DIR/bin/$EXEC
wget -c "https://github.com/jeffotoni/gobootcamp/blob/main/install/v1/gobootcamp" -P "$DIR"
echo "..."
sleep 1
chmod 755 $DIR/$EXEC

if [ -e "$ZSHRC" ] ; then
echo "\033[0;32m#########################################################\033[0m"
echo "atualizando ambiente $ZSHRC"
else
	if [ -e "$BASHRC" ] ; then
		echo "\033[0;32m#########################################################\033[0m"
		echo "atualizando ambiente $BASHRC"
		. $BASHRC
	else
	echo "\nNao encontrei bashrc!!"
	fi
fi

echo "\033[0;33m######### Thanks for Download ##########\033[0m"
echo "\033[0;33m gobootcamp \033[0m"
echo "......................................................"