cut -d '"' -f6 \
	| sort \
	 | ./funiq -c -i -I  -d 20 \
	 |sort -rn
