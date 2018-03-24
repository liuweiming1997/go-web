.PHONY: xo

xo:
	@./tools/generate-xo-models.sh

dump:
	@./tools/db.sh dump

restore:
	@./tools/db.sh restore

deploy:
	@./tools/deploy.sh deploy

localtest:
	@./tools/localtest.sh localtest