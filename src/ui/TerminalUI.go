package main

import (
	"log"

	"github.com/marcusolsson/tui-go"
)

var logo = `     _____ __ ____  ___   ______________  
    / ___// //_/\ \/ / | / / ____/_  __/  
    \__ \/ ,<    \  /  |/ / __/   / /     
   ___/ / /| |   / / /|  / /___  / /      
  /____/_/ |_|  /_/_/ |_/_____/ /_/     `

func main() {
	user := tui.NewTextEdit()
	user.SetFocused(true)
	user.SetText("ZhengYa")

	password := tui.NewEntry()
	password.SetEchoMode(tui.EchoModePassword)

	form := tui.NewGrid(0, 0)
	form.AppendRow(tui.NewLabel("用户"), tui.NewLabel("密码"))
	form.AppendRow(user, password)

	status := tui.NewStatusBar("已确认.")

	login := tui.NewButton("[登录]")
	login.OnActivated(func(b *tui.Button) {
		status.SetText("登陆成功")
	})

	register := tui.NewButton("[注册]")
	register.OnActivated(func(b *tui.Button) {
		status.SetText("注册你妹")
		user.SetText("都没账户，还项填写用户名？？？")
	})

	buttons := tui.NewHBox(
		tui.NewSpacer(),
		tui.NewPadder(1, 0, login),
		tui.NewPadder(1, 0, register),
	)

	window := tui.NewVBox(
		tui.NewPadder(10, 1, tui.NewLabel(logo)),
		tui.NewPadder(12, 0, tui.NewLabel("欢迎来到 SkyNet! 选择登录或者注册")),
		tui.NewPadder(1, 1, form),
		buttons,
	)
	window.SetBorder(true)

	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		window,
		tui.NewSpacer(),
	)
	content := tui.NewHBox(tui.NewSpacer(), wrapper, tui.NewSpacer())

	root := tui.NewVBox(
		content,
		status,
		tui.NewHBox(
			status,
		),
	)

	tui.DefaultFocusChain.Set(user, password, login, register)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
