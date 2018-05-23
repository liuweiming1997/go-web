.PHONY: xo

xo:
	@./shell/generate-xo-models.sh

dump:
	@./shell/db.sh dump

restore:
	@./shell/db.sh restore

deploy:
	@./shell/deploy.sh deploy

localtest:
	@./shell/localtest.sh localtest

getRemote:
	@./shell/util.sh getRemote

stopRemote:
	@./shell/util.sh stopRemote

logRemote:
	@./shell/util.sh logRemote