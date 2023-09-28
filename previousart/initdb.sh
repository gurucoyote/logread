sqlite3 $1  << EOF
create table logs 
(ip TEXT,
 timestamp TEXT,
 status_code TEXT,
 bytes_sent TEXT,
 request_method TEXT,
 request_url TEXT,
  request_protocol TEXT,
 referrer TEXT,
 user_agent TEXT);
DROP VIEW IF EXISTS requests; 
create view requests as
select substr(timestamp , 1, 11) as day,
 substr(timestamp, 13, 2) as hour,
 substr(timestamp, 16, 2) as min,  
*
from logs 
EOF 
