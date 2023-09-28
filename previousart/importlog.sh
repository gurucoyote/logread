# first arg: dbfile
# second arg: logfile 
# use zcat -f so we can also cat uncompressed files
# explicitly use the gzip installed version in this path
(echo "begin transaction;" ; \
/usr/bin/zcat -f $2 \
	| ./nginx2sql.sh; \
	echo "end transaction;") \
	| sqlite3 $1 
