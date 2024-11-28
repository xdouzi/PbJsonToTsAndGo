package main

import (
	"PbJsonToTsAndGo/Toos/json-gen-go/to_go"
	"PbJsonToTsAndGo/Toos/json-gen-java/to_java"
	"PbJsonToTsAndGo/Toos/json-gen-ts/to_ts"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// 创建应用
	myApp := app.New()
	myWindow := myApp.NewWindow("pb-gen-JsonToos ")
	// 设置窗口大小为 800x600
	myWindow.Resize(fyne.NewSize(800, 600))
	//---------------------------------------------------------------------------

	// 创建一个文本控件
	txt_1 := widget.NewLabel("通过Protol里的 *.proto 生成对应的Protocal") // 设置文本内容
	// 设置按钮的位置和大小
	txt_1.Resize(fyne.NewSize(200, 40)) // 宽高为 100x40
	txt_1.Move(fyne.NewPos(300, 10))    // 坐标 x=10, y=10

	// 创建一个文本控件
	txt_2 := widget.NewLabel("author:yh :技术支持联系408309839@qq.com") // 设置文本内容
	// 设置按钮的位置和大小
	txt_2.Resize(fyne.NewSize(200, 40)) // 宽高为 100x40
	txt_2.Move(fyne.NewPos(300, 30))    // 坐标 x=10, y=10

	// 创建一个文本控件
	txt_3 := widget.NewLabel("针对cocos 项目功能和web网站json传输,强转对应语言协议结构类型") // 设置文本内容
	// 设置按钮的位置和大小
	txt_3.Resize(fyne.NewSize(200, 40)) // 宽高为 100x40
	txt_3.Move(fyne.NewPos(300, 50))    // 坐标 x=10, y=10
	//---------------------------------------------------------------------------
	// 创建按钮
	but_ts := widget.NewButton("json-gen-ts", func() {
		println("Button clicked!")
		OnOpenDo(1)
	})
	// 设置按钮的位置和大小
	but_ts.Resize(fyne.NewSize(100, 40)) // 宽高为 100x40
	but_ts.Move(fyne.NewPos(10, 10))     // 坐标 x=10, y=10
	//---------------------------------------------------------------------------
	// 创建按钮
	but_go := widget.NewButton("json-gen-go", func() {
		println("Button clicked!")
		OnOpenDo(2)
	})
	// 设置按钮的位置和大小
	but_go.Resize(fyne.NewSize(100, 40)) // 宽高为 100x40
	but_go.Move(fyne.NewPos(10, 60))     // 坐标 x=10, y=10
	//---------------------------------------------------------------------------
	// 创建按钮
	but_java := widget.NewButton("json-gen-java", func() {
		println("Button clicked!")
		OnOpenDo(3)
	})
	// 设置按钮的位置和大小
	but_java.Resize(fyne.NewSize(100, 40)) // 宽高为 100x40
	but_java.Move(fyne.NewPos(10, 110))    // 坐标 x=10, y=10

	// 使用自由布局
	myWindow.SetContent(container.NewWithoutLayout(
		txt_1, txt_2, txt_3,
		but_ts,
		but_go,
		but_java,
	))

	// 显示窗口
	myWindow.ShowAndRun()
}

func OnOpenDo(index int32) {
	dirPath := "./Protol"

	switch index {
	case 1: //ts
		to_ts.NewPbToTsMain().Start(dirPath)
	case 2: //golang
		to_go.NewPbToGoMain().Start(dirPath)
	case 3: //java
		to_java.NewPbToJavaMain().Start(dirPath)
	default:

	}
}
