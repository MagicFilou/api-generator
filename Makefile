push: document
	@git add -A
	@git commit -m "$m"
	@git push

update:
	@grep 'module ' go.mod -v | grep 'git.wult.io/' | grep -v 'require'| awk '{print $$1}' | xargs go get -u 

document:
	./scripts/doc.sh
