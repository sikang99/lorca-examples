#
# Makefile for lorca-examples
#
all: usage
usage:
	@echo "make [edit|test||git]"

#------------------------------------------------------------------------------
edit e:
	@echo "make (edit) [make|readme]"

edit-make em:
	vi Makefile

edit-readme er:
	vi README.md

#------------------------------------------------------------------------------
test t:
	@echo "make (test) [object]"

test-object to:

#------------------------------------------------------------------------------
git g:
	@echo "make (git) [update|login|tag|status]"

git-update gu:
	git add .gitignore Makefile* *.md hello/ call/ counter/ gcal/ gui/ search/
	#git commit -m "initial commit"
	#git commit -m "add Makefile.dart-web"
	git commit -m "update commentaries"
	#git commit -m "rename /reference into /code"
	git push -u origin master 

git-remove gr:
	git rm -r VIM.md
	git commit -m "Rename VIM.md into TOOLS.md"
	git push origin master

git-login gl:
	git config --global user.email "sikang99@gmail.com"
	git config --global user.name "Stoney Kang"
	git config --global push.default matching
	#git config --global push.default simple
	git config credential.helper store

git-tag gt:
	git tag v0.0.3
	git push --tags

git-status gs:
	git diff
	git status
	git log --oneline -5
# ---------------------------------------------------------------------------

