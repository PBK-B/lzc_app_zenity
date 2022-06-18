// Code generated by 'go generate'; DO NOT EDIT.

package win

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modcomctl32 = windows.NewLazySystemDLL("comctl32.dll")
	modcomdlg32 = windows.NewLazySystemDLL("comdlg32.dll")
	modgdi32    = windows.NewLazySystemDLL("gdi32.dll")
	modntdll    = windows.NewLazySystemDLL("ntdll.dll")
	modole32    = windows.NewLazySystemDLL("ole32.dll")
	modshell32  = windows.NewLazySystemDLL("shell32.dll")
	moduser32   = windows.NewLazySystemDLL("user32.dll")
	modwtsapi32 = windows.NewLazySystemDLL("wtsapi32.dll")

	procInitCommonControlsEx        = modcomctl32.NewProc("InitCommonControlsEx")
	procChooseColorW                = modcomdlg32.NewProc("ChooseColorW")
	procCommDlgExtendedError        = modcomdlg32.NewProc("CommDlgExtendedError")
	procGetOpenFileNameW            = modcomdlg32.NewProc("GetOpenFileNameW")
	procGetSaveFileNameW            = modcomdlg32.NewProc("GetSaveFileNameW")
	procCreateFontIndirectW         = modgdi32.NewProc("CreateFontIndirectW")
	procDeleteObject                = modgdi32.NewProc("DeleteObject")
	procGetDeviceCaps               = modgdi32.NewProc("GetDeviceCaps")
	procRtlGetNtVersionNumbers      = modntdll.NewProc("RtlGetNtVersionNumbers")
	procCoCreateInstance            = modole32.NewProc("CoCreateInstance")
	procCoTaskMemFree               = modole32.NewProc("CoTaskMemFree")
	procSHBrowseForFolder           = modshell32.NewProc("SHBrowseForFolder")
	procSHCreateItemFromParsingName = modshell32.NewProc("SHCreateItemFromParsingName")
	procSHGetPathFromIDListEx       = modshell32.NewProc("SHGetPathFromIDListEx")
	procShell_NotifyIconW           = modshell32.NewProc("Shell_NotifyIconW")
	procGetDlgCtrlID                = moduser32.NewProc("GetDlgCtrlID")
	procWTSSendMessageW             = modwtsapi32.NewProc("WTSSendMessageW")
)

func InitCommonControlsEx(icc *INITCOMMONCONTROLSEX) (ok bool) {
	r0, _, _ := syscall.Syscall(procInitCommonControlsEx.Addr(), 1, uintptr(unsafe.Pointer(icc)), 0, 0)
	ok = r0 != 0
	return
}

func ChooseColor(cc *CHOOSECOLOR) (ok bool) {
	r0, _, _ := syscall.Syscall(procChooseColorW.Addr(), 1, uintptr(unsafe.Pointer(cc)), 0, 0)
	ok = r0 != 0
	return
}

func commDlgExtendedError() (code int) {
	r0, _, _ := syscall.Syscall(procCommDlgExtendedError.Addr(), 0, 0, 0, 0)
	code = int(r0)
	return
}

func GetOpenFileName(ofn *OPENFILENAME) (ok bool) {
	r0, _, _ := syscall.Syscall(procGetOpenFileNameW.Addr(), 1, uintptr(unsafe.Pointer(ofn)), 0, 0)
	ok = r0 != 0
	return
}

func GetSaveFileName(ofn *OPENFILENAME) (ok bool) {
	r0, _, _ := syscall.Syscall(procGetSaveFileNameW.Addr(), 1, uintptr(unsafe.Pointer(ofn)), 0, 0)
	ok = r0 != 0
	return
}

func CreateFontIndirect(lf *LOGFONT) (font Handle) {
	r0, _, _ := syscall.Syscall(procCreateFontIndirectW.Addr(), 1, uintptr(unsafe.Pointer(lf)), 0, 0)
	font = Handle(r0)
	return
}

func DeleteObject(o Handle) (ok bool) {
	r0, _, _ := syscall.Syscall(procDeleteObject.Addr(), 1, uintptr(o), 0, 0)
	ok = r0 != 0
	return
}

func GetDeviceCaps(dc Handle, index int) (cap int) {
	r0, _, _ := syscall.Syscall(procGetDeviceCaps.Addr(), 2, uintptr(dc), uintptr(index), 0)
	cap = int(r0)
	return
}

func RtlGetNtVersionNumbers(major *uint32, minor *uint32, build *uint32) {
	syscall.Syscall(procRtlGetNtVersionNumbers.Addr(), 3, uintptr(unsafe.Pointer(major)), uintptr(unsafe.Pointer(minor)), uintptr(unsafe.Pointer(build)))
	return
}

func CoCreateInstance(clsid uintptr, unkOuter unsafe.Pointer, clsContext int32, iid uintptr, address unsafe.Pointer) (res error) {
	r0, _, _ := syscall.Syscall6(procCoCreateInstance.Addr(), 5, uintptr(clsid), uintptr(unkOuter), uintptr(clsContext), uintptr(iid), uintptr(address), 0)
	if r0 != 0 {
		res = syscall.Errno(r0)
	}
	return
}

func CoTaskMemFree(address uintptr) {
	syscall.Syscall(procCoTaskMemFree.Addr(), 1, uintptr(address), 0, 0)
	return
}

func SHBrowseForFolder(bi *BROWSEINFO) (ptr uintptr) {
	r0, _, _ := syscall.Syscall(procSHBrowseForFolder.Addr(), 1, uintptr(unsafe.Pointer(bi)), 0, 0)
	ptr = uintptr(r0)
	return
}

func SHCreateItemFromParsingName(path *uint16, bc unsafe.Pointer, iid uintptr, item **IShellItem) (res error) {
	r0, _, _ := syscall.Syscall6(procSHCreateItemFromParsingName.Addr(), 4, uintptr(unsafe.Pointer(path)), uintptr(bc), uintptr(iid), uintptr(unsafe.Pointer(item)), 0, 0)
	if r0 != 0 {
		res = syscall.Errno(r0)
	}
	return
}

func SHGetPathFromIDListEx(ptr uintptr, path *uint16, pathLen int, opts int) (ok bool) {
	r0, _, _ := syscall.Syscall6(procSHGetPathFromIDListEx.Addr(), 4, uintptr(ptr), uintptr(unsafe.Pointer(path)), uintptr(pathLen), uintptr(opts), 0, 0)
	ok = r0 != 0
	return
}

func ShellNotifyIcon(message uint32, data *NOTIFYICONDATA) (ret int, err error) {
	r0, _, e1 := syscall.Syscall(procShell_NotifyIconW.Addr(), 2, uintptr(message), uintptr(unsafe.Pointer(data)), 0)
	ret = int(r0)
	if ret == 0 {
		err = errnoErr(e1)
	}
	return
}

func GetDlgCtrlID(wnd HWND) (ret int) {
	r0, _, _ := syscall.Syscall(procGetDlgCtrlID.Addr(), 1, uintptr(wnd), 0, 0)
	ret = int(r0)
	return
}

func WTSSendMessage(server Handle, sessionID uint32, title *uint16, titleLength int, message *uint16, messageLength int, style uint32, timeout int, response *uint32, wait bool) (err error) {
	var _p0 uint32
	if wait {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall12(procWTSSendMessageW.Addr(), 10, uintptr(server), uintptr(sessionID), uintptr(unsafe.Pointer(title)), uintptr(titleLength), uintptr(unsafe.Pointer(message)), uintptr(messageLength), uintptr(style), uintptr(timeout), uintptr(unsafe.Pointer(response)), uintptr(_p0), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}
