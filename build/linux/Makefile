# If PREFIX isn't provided, we check for $(DESTDIR)/usr/local and use that if it exists.
# Otherwice we fall back to using /usr.

LOCAL != test -d $(DESTDIR)/usr/local && echo -n "/local" || echo -n ""
LOCAL ?= $(shell test -d $(DESTDIR)/usr/local && echo "/local" || echo "")
PREFIX ?= /usr$(LOCAL)

Name := "toggl-redmine"
Exec := "toggl-redmine"
Icon := "toggl-redmine.png"

default:
	# User install
	# Run "make install" to install in ~/.local/
	# Run "make uninstall" to uninstall from ~/.local/

install:
	install -Dm00644 usr/local/share/applications/$(Name).desktop $(DESTDIR)$(HOME)/.local/share/applications/$(Name).desktop
	install -Dm00755 usr/local/bin/$(Exec) $(DESTDIR)$(HOME)/.local/bin/$(Exec)
	install -Dm00644 usr/local/share/pixmaps/$(Icon) $(DESTDIR)$(HOME)/.local/share/icons/$(Icon)
	sed -i -e "s,Exec=.*,Exec=$(DESTDIR)$(HOME)/.local/bin/$(Exec),g" $(DESTDIR)$(HOME)/.local/share/applications/$(Name).desktop

uninstall:
	-rm $(DESTDIR)$(HOME)/.local/share/applications/$(Name).desktop
	-rm $(DESTDIR)$(HOME)/.local/bin/$(Exec)
	-rm $(DESTDIR)$(HOME)/.local/share/icons/$(Icon)
