time: ts/time.go
	go build ts/time.go

list: list/list_timers.go
	go build list/list_timers.go

create: create/create_timers.go
	go build create/create_timers.go

install-time: time
	mv time /Users/wwei/Library/Application\ Support/Alfred/Alfred.alfredpreferences/workflows/user.workflow.448E5BE4-D892-47B9-9DDF-23A64ECCB8D2/workflow/time

install-list: list
	mv list_timers /Users/wwei/Library/Application\ Support/Alfred/Alfred.alfredpreferences/workflows/user.workflow.448E5BE4-D892-47B9-9DDF-23A64ECCB8D2/workflow/list_timers

install-create: create
	mv create_timers /Users/wwei/Library/Application\ Support/Alfred/Alfred.alfredpreferences/workflows/user.workflow.448E5BE4-D892-47B9-9DDF-23A64ECCB8D2/workflow/create_timers

clean:
	rm list_timers create_timers time

all: time list create

install-all: install-time install-list install-create
