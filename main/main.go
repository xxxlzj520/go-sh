package main

import (
	"fmt"
	sh "go-sh"
	"os/user"
	"strconv"
	"syscall"
)

func main() {
	user, err := user.Lookup("tteducom1111")
	if err != nil {
		fmt.Println(err)
		return
	}
	uid, _ := strconv.ParseUint(user.Uid, 10, 32)
	gid, _ := strconv.ParseUint(user.Gid, 10, 32)
	fmt.Println(uid, gid)
	//attr := exec.Cmd{}.SysProcAttr
	attr := &syscall.SysProcAttr{Credential: &syscall.Credential{
		Uid:         uint32(uid),
		Gid:         uint32(gid),
		NoSetGroups: true,
	}}
	//attr.Credential = &syscall.Credential{
	//	Uid:         uint32(uid),
	//	Gid:         uint32(gid),
	//	NoSetGroups: true,
	//}

	ccc, _ := execCmdShell("cd /root/1019;ln -s /home/tteducom1111 test;", attr)
	fmt.Println(ccc)
}
func execCmdShell(commond string, attr *syscall.SysProcAttr) (result string, err error) {
	session := sh.NewSession()
	session.ShowCMD = true
	session.Attr = attr
	session.Command("sh", "-c", commond)
	out, err1 := session.CombinedOutput() //.Output()
	if err1 != nil {
		fmt.Println(err1)
		err = err1
		//ErrorUp366(err1)
		return
	}
	result = string(out)
	fmt.Println(result)
	//InfoUp366(result)
	return
}
