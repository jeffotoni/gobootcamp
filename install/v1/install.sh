#!/bin/bash
##### @jeffotoni
DIR=/opt/gobootcamp/v1
DIR2=/opt/gobootcamp
EXEC=gobootcamp
ZSHRC=$HOME/.zshrc
BASHRC=$HOME/.bashrc

sudo rm -rf $DIR
sudo mkdir -p $DIR
sudo wget -c "https://github.com/jeffotoni/gobootcamp/blob/main/install/v1/gobootcamp?raw=true" -P "$DIR"
echo "..."
sleep 1
sudo chmod 755 -R $DIR2
sudo rm -f /usr/bin/$EXEC
sudo ln -s $DIR/$EXEC /usr/bin/$EXEC

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
echo "\033[0;33m You just need to type gobootcamp \033[0m"