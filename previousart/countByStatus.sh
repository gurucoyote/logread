 cut -d '"' -f3 \
| cut -d ' ' -f2 \
| sort | uniq -c | sort -rn
