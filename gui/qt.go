package gui

import (
	config2 "SsrMicroClient/config"
	"SsrMicroClient/net/delay"
	"SsrMicroClient/process/ServerControl"
	"SsrMicroClient/process/ssrcontrol"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"log"
	"os"
	"os/exec"
	"strconv"
)

type SsrMicroClientGUI struct {
	App                *widgets.QApplication
	MainWindow         *widgets.QMainWindow
	subscriptionWindow *widgets.QMainWindow
	settingWindow      *widgets.QMainWindow
	Session            *gui.QSessionManager
	ssrCmd             *exec.Cmd
	configPath         string
	settingConfig      *config2.Setting
	server             *ServerControl.ServerControl
}

func NewSsrMicroClientGUI(configPath string) (*SsrMicroClientGUI, error) {
	var err error
	microClientGUI := &SsrMicroClientGUI{}
	microClientGUI.configPath = configPath
	microClientGUI.settingConfig, err = config2.SettingDecodeJSON(microClientGUI.configPath)
	if err != nil {
		return microClientGUI, err
	}
	microClientGUI.ssrCmd = ssrcontrol.GetSsrCmd(microClientGUI.configPath)
	microClientGUI.App = widgets.NewQApplication(len(os.Args), os.Args)
	microClientGUI.App.SetApplicationName("SsrMicroClient")
	microClientGUI.App.SetQuitOnLastWindowClosed(false)
	microClientGUI.App.ConnectAboutToQuit(func() {
		if microClientGUI.ssrCmd.Process != nil {
			err = microClientGUI.ssrCmd.Process.Kill()
			if err != nil {
				//	do something
				log.Println(err)
			}
			_, err = microClientGUI.ssrCmd.Process.Wait()
			if err != nil {
				//	do something
				log.Println(err)
			}
		}
	})

	microClientGUI.server = &ServerControl.ServerControl{}

	microClientGUI.Session = gui.NewQSessionManagerFromPointer(nil)
	microClientGUI.App.SaveStateRequest(microClientGUI.Session)

	microClientGUI.MainWindow = widgets.NewQMainWindow(nil, 0)
	microClientGUI.createMainWindow()
	microClientGUI.subscriptionWindow = widgets.NewQMainWindow(microClientGUI.MainWindow, 0)
	microClientGUI.createSubscriptionWindow()
	microClientGUI.settingWindow = widgets.NewQMainWindow(microClientGUI.MainWindow, 0)
	microClientGUI.createSettingWindow()

	if microClientGUI.settingConfig.Bypass == true {
		microClientGUI.server.ServerStart()
	}

	return microClientGUI, nil
}

func (ssrMicroClientGUI *SsrMicroClientGUI) openMainWindow() {
	if ssrMicroClientGUI.MainWindow.IsHidden() == false {
		ssrMicroClientGUI.MainWindow.Hide()
	}
	ssrMicroClientGUI.MainWindow.Move2((ssrMicroClientGUI.App.Desktop().Width()-ssrMicroClientGUI.MainWindow.Width())/2, (ssrMicroClientGUI.App.Desktop().Height()-ssrMicroClientGUI.MainWindow.Height())/2)
	ssrMicroClientGUI.MainWindow.Show()
}

func (ssrMicroClientGUI *SsrMicroClientGUI) openSubscriptionWindow() {
	if ssrMicroClientGUI.subscriptionWindow.IsHidden() == false {
		ssrMicroClientGUI.subscriptionWindow.Close()
	}

	ssrMicroClientGUI.subscriptionWindow.Move2((ssrMicroClientGUI.App.Desktop().Width()-ssrMicroClientGUI.subscriptionWindow.Width())/2, (ssrMicroClientGUI.App.Desktop().Height()-ssrMicroClientGUI.subscriptionWindow.Height())/2)
	ssrMicroClientGUI.subscriptionWindow.Show()
}

func (ssrMicroClientGUI *SsrMicroClientGUI) openSettingWindow() {
	if ssrMicroClientGUI.settingWindow.IsHidden() == false {
		ssrMicroClientGUI.settingWindow.Close()
	}
	ssrMicroClientGUI.settingWindow.Move2((ssrMicroClientGUI.App.Desktop().Width()-ssrMicroClientGUI.settingWindow.Width())/2, (ssrMicroClientGUI.App.Desktop().Height()-ssrMicroClientGUI.settingWindow.Height())/2)
	ssrMicroClientGUI.settingWindow.Show()
}

func (ssrMicroClientGUI *SsrMicroClientGUI) createMainWindow() {
	ssrMicroClientGUI.MainWindow.SetFixedSize2(600, 400)
	ssrMicroClientGUI.MainWindow.SetWindowTitle("SsrMicroClient")
	icon := gui.NewQIcon5(ssrMicroClientGUI.configPath + "/SsrMicroClient.png")
	ssrMicroClientGUI.MainWindow.SetWindowIcon(icon)

	trayIcon := widgets.NewQSystemTrayIcon(ssrMicroClientGUI.MainWindow)
	trayIcon.SetIcon(icon)
	menu := widgets.NewQMenu(nil)
	ssrMicroClientTrayIconMenu := widgets.NewQAction2("SsrMicroClient", ssrMicroClientGUI.MainWindow)
	ssrMicroClientTrayIconMenu.ConnectTriggered(func(bool2 bool) {
		ssrMicroClientGUI.openMainWindow()
	})
	subscriptionTrayIconMenu := widgets.NewQAction2("subscription", ssrMicroClientGUI.MainWindow)
	subscriptionTrayIconMenu.ConnectTriggered(func(bool2 bool) {
		ssrMicroClientGUI.openSubscriptionWindow()
	})

	settingTrayIconMenu := widgets.NewQAction2("setting", ssrMicroClientGUI.MainWindow)
	settingTrayIconMenu.ConnectTriggered(func(bool2 bool) {
		ssrMicroClientGUI.openSettingWindow()
	})

	exit := widgets.NewQAction2("exit", ssrMicroClientGUI.MainWindow)
	exit.ConnectTriggered(func(bool2 bool) {
		ssrMicroClientGUI.App.Quit()
	})
	actions := []*widgets.QAction{ssrMicroClientTrayIconMenu,
		subscriptionTrayIconMenu, settingTrayIconMenu, exit}
	menu.AddActions(actions)
	trayIcon.SetContextMenu(menu)
	trayIcon.SetToolTip("")
	trayIcon.Show()

	statusLabel := widgets.NewQLabel2("status", ssrMicroClientGUI.MainWindow,
		core.Qt__WindowType(0x00000000))
	statusLabel.SetGeometry(core.NewQRect2(core.NewQPoint2(40, 10),
		core.NewQPoint2(130, 40)))
	statusLabel2 := widgets.NewQLabel2("", ssrMicroClientGUI.MainWindow,
		core.Qt__WindowType(0x00000000))
	statusLabel2.SetGeometry(core.NewQRect2(core.NewQPoint2(130, 10),
		core.NewQPoint2(560, 40)))

	nowNodeLabel := widgets.NewQLabel2("now node", ssrMicroClientGUI.MainWindow,
		core.Qt__WindowType(0x00000000))
	nowNodeLabel.SetGeometry(core.NewQRect2(core.NewQPoint2(40, 60),
		core.NewQPoint2(130, 90)))
	nowNode, err := config2.GetNowNode(ssrMicroClientGUI.configPath)
	if err != nil {
		ssrMicroClientGUI.MessageBox(err.Error())
		return
	}
	nowNodeLabel2 := widgets.NewQLabel2(nowNode["remarks"]+" - "+
		nowNode["group"], ssrMicroClientGUI.MainWindow, core.Qt__WindowType(0x00000000))
	nowNodeLabel2.SetGeometry(core.NewQRect2(core.NewQPoint2(130, 60),
		core.NewQPoint2(560, 90)))

	groupLabel := widgets.NewQLabel2("group", ssrMicroClientGUI.MainWindow,
		core.Qt__WindowType(0x00000000))
	groupLabel.SetGeometry(core.NewQRect2(core.NewQPoint2(40, 110),
		core.NewQPoint2(130, 140)))
	groupCombobox := widgets.NewQComboBox(ssrMicroClientGUI.MainWindow)
	group, err := config2.GetGroup(ssrMicroClientGUI.configPath)
	if err != nil {
		ssrMicroClientGUI.MessageBox(err.Error())
		return
	}
	groupCombobox.AddItems(group)
	groupCombobox.SetCurrentTextDefault(nowNode["group"])
	groupCombobox.SetGeometry(core.NewQRect2(core.NewQPoint2(130, 110),
		core.NewQPoint2(450, 140)))
	refreshButton := widgets.NewQPushButton2("refresh", ssrMicroClientGUI.MainWindow)
	refreshButton.SetGeometry(core.NewQRect2(core.NewQPoint2(460, 110),
		core.NewQPoint2(560, 140)))

	nodeLabel := widgets.NewQLabel2("node", ssrMicroClientGUI.MainWindow,
		core.Qt__WindowType(0x00000000))
	nodeLabel.SetGeometry(core.NewQRect2(core.NewQPoint2(40, 160),
		core.NewQPoint2(130, 190)))
	nodeCombobox := widgets.NewQComboBox(ssrMicroClientGUI.MainWindow)
	node, err := config2.GetNode(ssrMicroClientGUI.configPath, groupCombobox.CurrentText())
	if err != nil {
		ssrMicroClientGUI.MessageBox(err.Error())
		return
	}
	nodeCombobox.AddItems(node)
	nodeCombobox.SetCurrentTextDefault(nowNode["remarks"])
	nodeCombobox.SetGeometry(core.NewQRect2(core.NewQPoint2(130, 160),
		core.NewQPoint2(450, 190)))
	startButton := widgets.NewQPushButton2("start", ssrMicroClientGUI.MainWindow)

	// wait the last ssr process finished
	waitChan := make(chan bool, 0)
	go func() {
		waitChan <- true
	}()
	start := func() {
		if err := ssrMicroClientGUI.ssrCmd.Start(); err != nil {
			log.Println(err)
		}
		statusLabel2.SetText("<b><font color=green>running(pid:" + strconv.Itoa(ssrMicroClientGUI.ssrCmd.Process.Pid) + ")</font></b>")
		trayIcon.SetToolTip("running(pid:" + strconv.Itoa(ssrMicroClientGUI.ssrCmd.Process.Pid) + ")")
		if _, err := ssrMicroClientGUI.ssrCmd.Process.Wait(); err != nil {
			log.Println(err)
		}
		statusLabel2.SetText("<b><font color=red>stop</font></b>")
		trayIcon.SetToolTip("stop")
		waitChan <- true
	}
	startButton.ConnectClicked(func(bool2 bool) {
		group := groupCombobox.CurrentText()
		remarks := nodeCombobox.CurrentText()
		//_, exist := process.Get(ssrMicroClientGUI.configPath)
		log.Println(ssrMicroClientGUI.ssrCmd.Process, ssrMicroClientGUI.ssrCmd.ProcessState)
		if group == nowNode["group"] && remarks == nowNode["remarks"] && ssrMicroClientGUI.ssrCmd.Process != nil {
			if ssrMicroClientGUI.ssrCmd.Process.Pid != -1 {
				return
			}
		} else if group != nowNode["group"] || remarks != nowNode["remarks"] {
			err := config2.ChangeNowNode2(ssrMicroClientGUI.configPath, group, remarks)
			if err != nil {
				ssrMicroClientGUI.MessageBox(err.Error())
				return
			}
			nowNode, err = config2.GetNowNode(ssrMicroClientGUI.configPath)
			if err != nil {
				ssrMicroClientGUI.MessageBox(err.Error())
				return
			}
			nowNodeLabel2.SetText(nowNode["remarks"] + " - " + nowNode["group"])
			if ssrMicroClientGUI.ssrCmd.Process != nil {
				if err := ssrMicroClientGUI.ssrCmd.Process.Kill(); err != nil {
					log.Println(err)
				}
				ssrMicroClientGUI.ssrCmd = nil
			}
		}
		<-waitChan
		ssrMicroClientGUI.ssrCmd = ssrcontrol.GetSsrCmd(ssrMicroClientGUI.configPath)
		go func() {
			start()
		}()
	})

	startButton.SetGeometry(core.NewQRect2(core.NewQPoint2(460, 160),
		core.NewQPoint2(560, 190)))

	delayLabel := widgets.NewQLabel2("delay", ssrMicroClientGUI.MainWindow,
		core.Qt__WindowType(0x00000000))
	delayLabel.SetGeometry(core.NewQRect2(core.NewQPoint2(40, 210),
		core.NewQPoint2(130, 240)))
	delayLabel2 := widgets.NewQLabel2("", ssrMicroClientGUI.MainWindow,
		core.Qt__WindowType(0x00000000))
	delayLabel2.SetGeometry(core.NewQRect2(core.NewQPoint2(130, 210),
		core.NewQPoint2(450, 240)))
	delayButton := widgets.NewQPushButton2("get delay", ssrMicroClientGUI.MainWindow)
	delayButton.ConnectClicked(func(bool2 bool) {
		go func() {
			group := groupCombobox.CurrentText()
			remarks := nodeCombobox.CurrentText()
			node, err := config2.GetOneNode(ssrMicroClientGUI.configPath, group, remarks)
			if err != nil {
				ssrMicroClientGUI.MessageBox(err.Error())
				return
			}
			delayTmp, isSuccess, err := delay.TCPDelay(node.Server,
				node.ServerPort)
			var delayString string
			if err != nil {
				//ssrMicroClientGUI.MessageBox(err.Error())
				delayString = err.Error()
			} else {
				delayString = delayTmp.String()
			}
			if isSuccess == false {
				delayString = "delay > 3s or server can not connect"
			}
			delayLabel2.SetText(delayString)
		}()
	})
	delayButton.SetGeometry(core.NewQRect2(core.NewQPoint2(460, 210),
		core.NewQPoint2(560, 240)))

	groupCombobox.ConnectCurrentTextChanged(func(string2 string) {
		node, err := config2.GetNode(ssrMicroClientGUI.configPath,
			groupCombobox.CurrentText())
		if err != nil {
			ssrMicroClientGUI.MessageBox(err.Error())
		}
		nodeCombobox.Clear()
		nodeCombobox.AddItems(node)
	})

	subButton := widgets.NewQPushButton2("subscription setting", ssrMicroClientGUI.MainWindow)
	subButton.SetGeometry(core.NewQRect2(core.NewQPoint2(40, 260),
		core.NewQPoint2(290, 290)))
	subButton.ConnectClicked(func(bool2 bool) {
		ssrMicroClientGUI.openSubscriptionWindow()
	})

	subUpdateButton := widgets.NewQPushButton2("subscription Update", ssrMicroClientGUI.MainWindow)
	subUpdateButton.SetGeometry(core.NewQRect2(core.NewQPoint2(300, 260),
		core.NewQPoint2(560, 290)))
	subUpdateButton.ConnectClicked(func(bool2 bool) {
		message := widgets.NewQMessageBox(ssrMicroClientGUI.MainWindow)
		message.SetText("Updating!")
		message.Show()
		ch := make(chan bool)
		go func() {
			if err := config2.SsrJSON(ssrMicroClientGUI.configPath); err != nil {
				ssrMicroClientGUI.MessageBox(err.Error())
			}
			ch <- true
		}()
		<-ch
		message.SetText("Updated!")
		group, err = config2.GetGroup(ssrMicroClientGUI.configPath)
		if err != nil {
			ssrMicroClientGUI.MessageBox(err.Error())
			return
		}
		groupCombobox.Clear()
		groupCombobox.AddItems(group)
		groupCombobox.SetCurrentText(nowNode["group"])
		node, err = config2.GetNode(ssrMicroClientGUI.configPath, groupCombobox.CurrentText())
		if err != nil {
			ssrMicroClientGUI.MessageBox(err.Error())
			return
		}
		nodeCombobox.Clear()
		nodeCombobox.AddItems(node)
		nodeCombobox.SetCurrentText(nowNode["remarks"])
	})

	settingButton := widgets.NewQPushButton2("Setting", ssrMicroClientGUI.MainWindow)
	settingButton.SetGeometry(core.NewQRect2(core.NewQPoint2(40, 300),
		core.NewQPoint2(290, 330)))
	settingButton.ConnectClicked(func(bool2 bool) {
		ssrMicroClientGUI.openSettingWindow()
	})

	if ssrMicroClientGUI.settingConfig.AutoStartSsr == true {
		if ssrMicroClientGUI.ssrCmd.Process != nil {
			if ssrMicroClientGUI.ssrCmd.Process.Pid != -1 {
				return
			}
		}
		startButton.Click()
	}
}

func (ssrMicroClientGUI *SsrMicroClientGUI) createSubscriptionWindow() {
	ssrMicroClientGUI.subscriptionWindow.SetFixedSize2(700, 100)
	ssrMicroClientGUI.subscriptionWindow.SetWindowTitle("subscription")
	ssrMicroClientGUI.subscriptionWindow.ConnectCloseEvent(func(event *gui.QCloseEvent) {
		event.Ignore()
		ssrMicroClientGUI.subscriptionWindow.Hide()
	})

	subLabel := widgets.NewQLabel2("subscription", ssrMicroClientGUI.subscriptionWindow,
		core.Qt__WindowType(0x00000000))
	subLabel.SetGeometry(core.NewQRect2(core.NewQPoint2(10, 10),
		core.NewQPoint2(130, 40)))
	subCombobox := widgets.NewQComboBox(ssrMicroClientGUI.subscriptionWindow)
	var link []string
	subRefresh := func() {
		subCombobox.Clear()
		var err error
		link, err = config2.GetLink(ssrMicroClientGUI.configPath)
		if err != nil {
			ssrMicroClientGUI.MessageBox(err.Error())
		}
		subCombobox.AddItems(link)
	}
	subRefresh()
	subCombobox.SetGeometry(core.NewQRect2(core.NewQPoint2(115, 10),
		core.NewQPoint2(600, 40)))

	deleteButton := widgets.NewQPushButton2("delete", ssrMicroClientGUI.subscriptionWindow)
	deleteButton.ConnectClicked(func(bool2 bool) {
		linkToDelete := subCombobox.CurrentText()
		if err := config2.RemoveLinkJSON2(linkToDelete,
			ssrMicroClientGUI.configPath); err != nil {
			ssrMicroClientGUI.MessageBox(err.Error())
		}
		subRefresh()
	})
	deleteButton.SetGeometry(core.NewQRect2(core.NewQPoint2(610, 10),
		core.NewQPoint2(690, 40)))

	lineText := widgets.NewQLineEdit(ssrMicroClientGUI.subscriptionWindow)
	lineText.SetGeometry(core.NewQRect2(core.NewQPoint2(115, 50),
		core.NewQPoint2(600, 80)))

	addButton := widgets.NewQPushButton2("add", ssrMicroClientGUI.subscriptionWindow)
	addButton.ConnectClicked(func(bool2 bool) {
		linkToAdd := lineText.Text()
		if linkToAdd == "" {
			return
		}
		for _, linkExisted := range link {
			if linkExisted == linkToAdd {
				return
			}
		}
		if err := config2.AddLinkJSON2(linkToAdd, ssrMicroClientGUI.configPath); err != nil {
			//log.Println(err)
			ssrMicroClientGUI.MessageBox(err.Error())
			return
		}
		subRefresh()
	})
	addButton.SetGeometry(core.NewQRect2(core.NewQPoint2(610, 50),
		core.NewQPoint2(690, 80)))

	ssrMicroClientGUI.subscriptionWindow.ConnectCloseEvent(func(event *gui.QCloseEvent) {
		ssrMicroClientGUI.subscriptionWindow.Close()
	})
}

func (ssrMicroClientGUI *SsrMicroClientGUI) createSettingWindow() {
	ssrMicroClientGUI.settingWindow.SetFixedSize2(430, 330)
	ssrMicroClientGUI.settingWindow.SetWindowTitle("setting")
	ssrMicroClientGUI.settingWindow.ConnectCloseEvent(func(event *gui.QCloseEvent) {
		event.Ignore()
		ssrMicroClientGUI.settingWindow.Hide()
	})

	autoStartSsr := widgets.NewQCheckBox2("auto Start ssr", ssrMicroClientGUI.settingWindow)
	autoStartSsr.SetChecked(ssrMicroClientGUI.settingConfig.AutoStartSsr)
	autoStartSsr.SetGeometry(core.NewQRect2(core.NewQPoint2(10, 0),
		core.NewQPoint2(490, 30)))

	httpProxyCheckBox := widgets.NewQCheckBox2("http proxy", ssrMicroClientGUI.settingWindow)
	httpProxyCheckBox.SetDisabled(true)
	httpProxyCheckBox.SetChecked(ssrMicroClientGUI.settingConfig.HttpProxy)
	httpProxyCheckBox.SetGeometry(core.NewQRect2(core.NewQPoint2(10, 40),
		core.NewQPoint2(130, 70)))

	bypassCheckBox := widgets.NewQCheckBox2("bypass",
		ssrMicroClientGUI.settingWindow)
	bypassCheckBox.SetDisabled(true)
	bypassCheckBox.SetChecked(ssrMicroClientGUI.settingConfig.Bypass)
	bypassCheckBox.SetGeometry(core.NewQRect2(core.NewQPoint2(140, 40),
		core.NewQPoint2(220, 70)))

	DnsOverHttpsCheckBox := widgets.NewQCheckBox2("Use DNSOverHTTPS",
		ssrMicroClientGUI.settingWindow)
	DnsOverHttpsCheckBox.SetChecked(ssrMicroClientGUI.settingConfig.IsDNSOverHTTPS)
	DnsOverHttpsCheckBox.SetGeometry(core.NewQRect2(core.NewQPoint2(230, 40),
		core.NewQPoint2(450, 70)))

	//httpBypassCheckBox := widgets.NewQCheckBox2("http bypass", ssrMicroClientGUI.settingWindow)
	//httpBypassCheckBox.SetChecked(ssrMicroClientGUI.settingConfig.HttpWithBypass)
	//httpBypassCheckBox.SetGeometry(core.NewQRect2(core.NewQPoint2(310, 40),
	//	core.NewQPoint2(450, 70)))

	//localAddressLabel := widgets.NewQLabel2("SSRAddress", ssrMicroClientGUI.settingWindow, 0)
	//localAddressLabel.SetGeometry(core.NewQRect2(core.NewQPoint2(10, 80),
	//	core.NewQPoint2(80, 110)))
	//localAddressLineText := widgets.NewQLineEdit(ssrMicroClientGUI.settingWindow)
	//localAddressLineText.SetText(ssrMicroClientGUI.settingConfig.LocalAddress)
	//localAddressLineText.SetGeometry(core.NewQRect2(core.NewQPoint2(90, 80),
	//	core.NewQPoint2(200, 110)))
	//
	//localPortLabel := widgets.NewQLabel2("SSRPort", ssrMicroClientGUI.settingWindow, 0)
	//localPortLabel.SetGeometry(core.NewQRect2(core.NewQPoint2(230, 80),
	//	core.NewQPoint2(300, 110)))
	//localPortLineText := widgets.NewQLineEdit(ssrMicroClientGUI.settingWindow)
	//localPortLineText.SetText(ssrMicroClientGUI.settingConfig.LocalPort)
	//localPortLineText.SetGeometry(core.NewQRect2(core.NewQPoint2(310, 80),
	//	core.NewQPoint2(420, 110)))

	httpAddressLabel := widgets.NewQLabel2("http", ssrMicroClientGUI.settingWindow, 0)
	httpAddressLabel.SetGeometry(core.NewQRect2(core.NewQPoint2(10, 120),
		core.NewQPoint2(70, 150)))
	httpAddressLineText := widgets.NewQLineEdit(ssrMicroClientGUI.settingWindow)
	httpAddressLineText.SetText(ssrMicroClientGUI.settingConfig.HttpProxyAddressAndPort)
	httpAddressLineText.SetGeometry(core.NewQRect2(core.NewQPoint2(80, 120),
		core.NewQPoint2(210, 150)))

	socks5BypassAddressLabel := widgets.NewQLabel2("socks5",
		ssrMicroClientGUI.settingWindow, 0)
	socks5BypassAddressLabel.SetGeometry(core.
		NewQRect2(core.NewQPoint2(220, 120), core.NewQPoint2(290, 150)))
	socks5BypassLineText := widgets.NewQLineEdit(ssrMicroClientGUI.settingWindow)
	socks5BypassLineText.SetText(ssrMicroClientGUI.settingConfig.Socks5WithBypassAddressAndPort)
	socks5BypassLineText.SetGeometry(core.NewQRect2(core.
		NewQPoint2(300, 120), core.NewQPoint2(420, 150)))

	dnsServerLabel := widgets.NewQLabel2("DNS", ssrMicroClientGUI.settingWindow, 0)
	dnsServerLabel.SetGeometry(core.NewQRect2(core.
		NewQPoint2(10, 160), core.NewQPoint2(100, 190)))
	dnsServerLineText := widgets.NewQLineEdit(ssrMicroClientGUI.settingWindow)
	dnsServerLineText.SetText(ssrMicroClientGUI.settingConfig.DnsServer)
	dnsServerLineText.SetGeometry(core.NewQRect2(core.
		NewQPoint2(110, 160), core.NewQPoint2(420, 190)))

	ssrPathLabel := widgets.NewQLabel2("ssrPath", ssrMicroClientGUI.settingWindow, 0)
	ssrPathLabel.SetGeometry(core.NewQRect2(core.
		NewQPoint2(10, 200), core.NewQPoint2(100, 230)))
	ssrPathLineText := widgets.NewQLineEdit(ssrMicroClientGUI.settingWindow)
	ssrPathLineText.SetText(ssrMicroClientGUI.settingConfig.SsrPath)
	ssrPathLineText.SetGeometry(core.NewQRect2(core.
		NewQPoint2(110, 200), core.NewQPoint2(420, 230)))

	BypassFileLabel := widgets.NewQLabel2("bypassFilePath", ssrMicroClientGUI.settingWindow, 0)
	BypassFileLabel.SetGeometry(core.NewQRect2(core.NewQPoint2(10, 240),
		core.NewQPoint2(100, 270)))
	BypassFileLineText := widgets.NewQLineEdit(ssrMicroClientGUI.settingWindow)
	BypassFileLineText.SetText(ssrMicroClientGUI.settingConfig.BypassFile)
	BypassFileLineText.SetGeometry(core.NewQRect2(core.NewQPoint2(110, 240),
		core.NewQPoint2(420, 270)))

	applyButton := widgets.NewQPushButton2("apply", ssrMicroClientGUI.settingWindow)
	applyButton.ConnectClicked(func(bool2 bool) {
		if socks5BypassLineText.Text() == "127.0.0.1:1083" || socks5BypassLineText.Text() == "0.0.0.0:1083" {
			ssrMicroClientGUI.MessageBox("You cant set the socks5 port to 1083,Please it.")
			return
		}
		ssrMicroClientGUI.settingConfig.AutoStartSsr = autoStartSsr.IsChecked()
		ssrMicroClientGUI.settingConfig.HttpProxy = httpProxyCheckBox.IsChecked()
		ssrMicroClientGUI.settingConfig.Bypass = bypassCheckBox.IsChecked()
		ssrMicroClientGUI.settingConfig.IsDNSOverHTTPS = DnsOverHttpsCheckBox.IsChecked()
		//ssrMicroClientGUI.settingConfig.HttpWithBypass = httpBypassCheckBox.IsChecked()
		//ssrMicroClientGUI.settingConfig.LocalAddress = localAddressLineText.Text()
		//ssrMicroClientGUI.settingConfig.LocalPort = localPortLineText.Text()
		//ssrMicroClientGUI.settingConfig.PythonPath = pythonPathLineText.Text()
		ssrMicroClientGUI.settingConfig.SsrPath = ssrPathLineText.Text()
		ssrMicroClientGUI.settingConfig.BypassFile = BypassFileLineText.Text()
		ssrMicroClientGUI.settingConfig.HttpProxyAddressAndPort = httpAddressLineText.Text()
		ssrMicroClientGUI.settingConfig.Socks5WithBypassAddressAndPort = socks5BypassLineText.Text()
		ssrMicroClientGUI.settingConfig.DnsServer = dnsServerLineText.Text()

		if err := config2.SettingEnCodeJSON(ssrMicroClientGUI.configPath, ssrMicroClientGUI.settingConfig); err != nil {
			//log.Println(err)
			ssrMicroClientGUI.MessageBox(err.Error())
		}
		ssrMicroClientGUI.server.ServerRestart()
	})
	applyButton.SetGeometry(core.NewQRect2(core.NewQPoint2(10, 280),
		core.NewQPoint2(90, 310)))

	ssrMicroClientGUI.settingWindow.ConnectCloseEvent(func(event *gui.QCloseEvent) {
		ssrMicroClientGUI.settingWindow.Close()
	})
}

func (ssrMicroClientGUI *SsrMicroClientGUI) MessageBox(text string) {
	message := widgets.NewQMessageBox(nil)
	message.SetText(text)
	message.Exec()
}
