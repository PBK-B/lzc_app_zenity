package zenity

import (
	"math/rand"
	"runtime"
	"syscall"
	"time"
	"unsafe"

	"github.com/ncruces/zenity/internal/win"
	"github.com/ncruces/zenity/internal/zenutil"
)

func notify(text string, opts options) error {
	if opts.ctx != nil && opts.ctx.Err() != nil {
		return opts.ctx.Err()
	}

	var args win.NOTIFYICONDATA
	args.StructSize = uint32(unsafe.Sizeof(args))
	args.ID = rand.Uint32()
	args.Flags = 0x00000010 // NIF_INFO
	args.State = 0x00000001 // NIS_HIDDEN

	info := syscall.StringToUTF16(text)
	copy(args.Info[:len(args.Info)-1], info)

	if opts.title != nil {
		title := syscall.StringToUTF16(*opts.title)
		copy(args.InfoTitle[:len(args.InfoTitle)-1], title)
	}

	switch opts.icon {
	case InfoIcon, QuestionIcon:
		args.InfoFlags |= 0x1 // NIIF_INFO
	case WarningIcon:
		args.InfoFlags |= 0x2 // NIIF_WARNING
	case ErrorIcon:
		args.InfoFlags |= 0x3 // NIIF_ERROR
	default:
		icon := getIcon(opts.icon)
		if icon.handle != 0 {
			defer icon.delete()
			args.Icon = win.Handle(icon.handle)
			args.Flags |= 0x00000002 // NIF_ICON
			args.InfoFlags |= 0x4    // NIIF_USER
		}
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	s, err := win.ShellNotifyIcon(win.NIM_ADD, &args)
	if s == 0 {
		if errno, ok := err.(syscall.Errno); ok && errno == 0 {
			return wtsMessage(text, opts)
		}
		return err
	}

	var major, minor, build uint32
	win.RtlGetNtVersionNumbers(&major, &minor, &build)

	// On Windows 7 (6.1) and lower, wait up to 10 seconds to clean up.
	if major < 6 || major == 6 && minor < 2 {
		if opts.ctx != nil {
			select {
			case <-opts.ctx.Done():
			case <-time.After(10 * time.Second):
			}
		} else {
			time.Sleep(10 * time.Second)
		}
	}

	win.ShellNotifyIcon(win.NIM_DELETE, &args)
	return nil
}

func wtsMessage(text string, opts options) error {
	var flags uint32

	switch opts.icon {
	case ErrorIcon:
		flags |= win.MB_ICONERROR
	case QuestionIcon:
		flags |= win.MB_ICONQUESTION
	case WarningIcon:
		flags |= win.MB_ICONWARNING
	case InfoIcon:
		flags |= win.MB_ICONINFORMATION
	}

	title := opts.title
	if title == nil {
		title = stringPtr("Notification")
	}

	timeout := zenutil.Timeout
	if timeout == 0 {
		timeout = 10
	}

	ptext := syscall.StringToUTF16(text)
	ptitle := syscall.StringToUTF16(*title)

	var res uint32
	return win.WTSSendMessage(
		win.WTS_CURRENT_SERVER_HANDLE, win.WTS_CURRENT_SESSION,
		&ptitle[0], 2*len(ptitle), &ptext[0], 2*len(ptext),
		flags, timeout, &res, false)
}
