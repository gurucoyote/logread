# parsing every line
gawk '{print  $1  " " substr($4, 2, length($4) - 8) " " $6 " " $7 " " substr($5, 2, length($5) -2)  " " substr($8, 2, length($8)-2) " "  $9; system("")}' FPAT='[^ ]*|"[^"]*"|\\[[^]]*\\]' | \
	# write each parse line to a SQL INSERT statement
(while read ip timestamp status_code bytes_sent request_method request_url request_protocol referrer user_agent; do 
echo  \
	"insert into logs (ip, timestamp, status_code, bytes_sent, request_method, request_url, request_protocol, referrer, user_agent) \
	values (\"$ip\",\"$timestamp\", \"$status_code\", \"$bytes_sent\", \"$request_method\", \"$request_url\", \"$request_protocol\" ,\"$referrer\" , $user_agent);"; \
done )
# code and ideas from here:
# https://funnybretzel.com/self-hosted-analytics-using-sqlite-and-metabase/
