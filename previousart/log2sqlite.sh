
# code and ideas from here:
# https://funnybretzel.com/self-hosted-analytics-using-sqlite-and-metabase/

# tail -f /var/log/nginx/access.log | \
# parsing every line
gawk '{print  $1  " " substr($4, 2, length($4) - 8) " " $6 " " $7 " " substr($5, 2, length($5) -2)  " " substr($8, 2, length($8)-2) " "  $9; system("")}' FPAT='[^ ]*|"[^"]*"|\\[[^]]*\\]' | \
# inserting every new line into sqlite
(while read ip timestamp status_code bytes_sent request_method request_url request_protocol referrer user_agent; do 
# sqlite3 -batch logs.db \
echo  \
	"insert into todo (ip, timestamp, status_code, bytes_sent, request_method, request_url, request_protocol, referrer, user_agent) \
	values (\"$ip\",\"$timestamp\", \"$status_code\", \"$bytes_sent\", \"$request_method\", \"$request_url\", \"$request_protocol\" ,\"$referrer\" , $user_agent);" \
>> logs.sql; 
done )
