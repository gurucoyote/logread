#!/bin/sh -e

rsync -vrc  \
	root@85.214.148.98:/DeinDokument/nginx/etc \
	./
rsync -vrc  \
	root@85.214.148.98:/DeinDokument/nginx/secrets-prod \
	./
