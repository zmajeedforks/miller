# Respective MANPATH entries would include /usr/local/share/man or $HOME/man.
# This should be run after make in the base directory but before make in the ../docs directory,
# since ../mlr is used to autogenerate ./manpage.txt which is used in ../docs.
# See also ../Makefile and
# https://miller.readthedocs.io/en/latest/build.html#creating-a-new-release-for-developers
PREFIX=/usr/local
INSTALLDIR=$(PREFIX)/share/man/man1

# This is normally done only on a development host. Through CI and
# package-installer, mlr.1 should be treated as an already-built artifact,
# needing only to be copied to its install dir.
# Copying to man1/mlr1 is so that you can then put the Miller dir in your
# $MANPATH and do "man mlr".
build:
	echo mkman start
	./mkman.rb > mlr.1
	./mkman.rb | groff -man -Tascii | col -b | expand -8 > manpage.txt
	cp manpage.txt ../docs/src/
	echo mkman end
	mkdir -p ./man1
	cp mlr.1 ./man1

# DESTDIR is for package installs; nominally blank when this is run interactively.
# See also https://www.gnu.org/prep/standards/html_node/DESTDIR.html
install:
	mkdir -p ./man1
	cp mlr.1 ./man1
	mkdir -p $(DESTDIR)/$(INSTALLDIR)
	cp mlr.1 $(DESTDIR)/$(INSTALLDIR)/mlr.1

.PHONY: build install
