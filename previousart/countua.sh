cut -d '"' -f6 \
| sort | uniq -c \
| sort -nr
