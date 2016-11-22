export UDIR= .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G   = et2g
export EG2GO  = eg2go

export GOARCH = amd64
export TARGET_ARCH = x86_64
export GOETHOSINCLUDE=/usr/lib64/go/pkg/ethos_$(GOARCH)
export GOLINUXINCLUDE=/usr/lib64/go/pkg/linux_$(GOARCH)


install.rootfs = /var/lib/ethos/ethos-default-$(TARGET_ARCH)/rootfs
install.minimaltd.rootfs = /var/lib/ethos/minimaltd/rootfs


.PHONY: all install
all: Box

install: Box
	ethosTypeInstall $(install.rootfs) $(install.minimaltd.rootfs) myBox
	install Box $(install.rootfs)/programs
	install Box $(install.minimaltd.rootfs)/programs
	echo -n /programs/Box | ethosStringEncode > $(install.rootfs)/etc/init/console
	mkdir -p $(install.rootfs)/user/nobody/myDir
	setfattr -n user.ethos.typeHash -v $(shell egPrint myBox hash Box) $(install.rootfs)/user/nobody/myDir
	

myBox.go: myBox.t
	$(ETN2GO) . myBox main $^

Box: Box.go myBox.go
	ethosGo $^ 

clean:
	rm -rf myBox/ myBoxIndex/
	rm -f myBox.go
	rm -f Box
	rm -f Box.goo.ethos
