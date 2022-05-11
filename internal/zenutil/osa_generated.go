// Code generated by zenity; DO NOT EDIT.
//go:build darwin

package zenutil

import (
	"encoding/json"
	"text/template"
)

var scripts = template.Must(template.New("").Funcs(template.FuncMap{"json": func(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}}).Parse(`
{{define "color" -}}
var app=Application.currentApplication()
app.includeStandardAdditions=true
app.activate()
ObjC.import('stdio')
ObjC.import('stdlib')
try{var res=app.chooseColor({defaultColor:{{json .}}})}catch(e){if(e.errorNumber===-128)$.exit(1)
$.dprintf(2,e)
$.exit(-1)}
{'rgb('+res.map(x=>Math.round(x*255))+')'}
{{- end}}
{{define "date" -}}
var app=Application.currentApplication()
app.includeStandardAdditions=true
app.activate()
ObjC.import('Cocoa')
ObjC.import('stdio')
ObjC.import('stdlib')
var date=$.NSDatePicker.alloc.init
date.setDatePickerStyle($.NSDatePickerStyleClockAndCalendar)
date.setDatePickerElements($.NSDatePickerElementFlagYearMonthDay)
date.setFrameSize(date.fittingSize)
{{- if .Date}}
date.setDateValue($.NSDate.dateWithTimeIntervalSince1970({{.Date}}))
{{- end}}
var alert=$.NSAlert.alloc.init
alert.setAccessoryView(date)
alert.setMessageText({{json .Text}})
alert.addButtonWithTitle({{json .OK}})
alert.addButtonWithTitle({{json .Cancel}}).keyEquivalent='\033'
{{- if .Info}}
alert.setInformativeText({{json .Info}})
{{- end}}
{{- if .Extra}}
alert.addButtonWithTitle({{json .Extra}})
{{- end}}
var res=alert.runModal
switch(res){case $.NSAlertThirdButtonReturn:$.puts({{json .Extra}})
case $.NSAlertSecondButtonReturn:$.exit(1)}
var fmt=$.NSDateFormatter.alloc.init
fmt.locale=$.NSLocale.localeWithLocaleIdentifier('en_US_POSIX')
fmt.dateFormat={{json .Format}}
fmt.stringFromDate(date.dateValue)
{{- end}}
{{define "dialog" -}}
var app=Application.currentApplication()
app.includeStandardAdditions=true
app.activate()
ObjC.import('stdio')
ObjC.import('stdlib')
var opts={{json .Options}}
{{- if .IconPath}}
opts.withIcon=Path({{json .IconPath}})
{{- end}}
try{var res=app.{{.Operation}}({{json .Text}},opts)}catch(e){if(e.errorNumber===-128)$.exit(1)
$.dprintf(2,e)
$.exit(-1)}
if(res.gaveUp){$.exit(5)}
if(res.buttonReturned==={{json .Extra}}){$.puts({{json .Extra}})
$.exit(1)}
res.textReturned
{{- end}}
{{define "file" -}}
var app=Application.currentApplication()
app.includeStandardAdditions=true
app.activate()
ObjC.import('stdio')
ObjC.import('stdlib')
try{var res=app.{{.Operation}}({{json .Options}})}catch(e){if(e.errorNumber===-128)$.exit(1)
$.dprintf(2,e)
$.exit(-1)}
if(Array.isArray(res)){res.join({{json .Separator}})}else{res.toString()}
{{- end}}
{{define "list" -}}
var app=Application.currentApplication()
app.includeStandardAdditions=true
ObjC.import('stdio')
ObjC.import('stdlib')
try{var res=app.chooseFromList({{json .Items}},{{json .Options}})}catch(e){$.dprintf(2,e)
$.exit(-1)}
if(res===false)$.exit(1)
if(res.length!==0)res.join({{json .Separator}})
{{- end}}
{{define "notify" -}}
var app=Application.currentApplication()
app.includeStandardAdditions=true
void app.displayNotification({{json .Text}},{{json .Options}})
{{- end}}
{{define "progress" -}}
var app=Application.currentApplication()
app.includeStandardAdditions=true
app.activate()
ObjC.import('stdlib')
ObjC.import('readline')
{{- if .Total}}
Progress.totalUnitCount={{.Total}}
{{- end}}
{{- if .Description}}
Progress.description={{json .Description}}
{{- end}}
while(true){try{var s=$.readline('')}catch(e){if(e.errorNumber===-128)$.exit(1)
break}
if(s.indexOf('#')===0){Progress.additionalDescription=s.slice(1)
continue}
var i=parseInt(s)
if(i>=0&&Progress.totalUnitCount>0){Progress.completedUnitCount=i}}
{{- end}}
{{define "pwd" -}}
var app=Application.currentApplication()
app.includeStandardAdditions=true
app.activate()
ObjC.import('stdio')
ObjC.import('stdlib')
var opts={{json .Options}}
{{- if .IconPath}}
opts.withIcon=Path({{json .IconPath}})
{{- end}}
function dialog(text){try{var res=app.displayDialog(text,opts)}catch(e){if(e.errorNumber===-128)$.exit(1)
$.dprintf(2,e)
$.exit(-1)}
if(res.gaveUp){$.exit(5)}
if(res.buttonReturned==={{json .Extra}}){$.puts({{json .Extra}})
$.exit(1)}
return res.textReturned}
var start=Date.now()
opts.defaultAnswer=''
var username=dialog('Username:')
{{- if .Options.Timeout}}
opts.givingUpAfter-=(Date.now()-start)/1000|0
{{- end}}
opts.hiddenAnswer=true
var password=dialog('Password:')
username+{{json .Separator}}+password
{{- end}}`))
