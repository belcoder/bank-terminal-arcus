#!/bin/sh

os="";

grep "debian" /etc/issue -i -q
if [ $? = '0' ]; then
	/home/reg/arcus/ipservice/setup_debian.sh
	exit
fi


ln -s /home/reg/arcus/ipservice/arcusip /etc/init.d/arcusip
chmod -v +x /etc/init.d/arcusip
chkconfig --add arcusip
