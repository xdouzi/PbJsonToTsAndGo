package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// 创建应用
	myApp := app.New()
	myWindow := myApp.NewWindow("Button Example")
	// 设置窗口大小为 800x600
	myWindow.Resize(fyne.NewSize(800, 600))
	//---------------------------------------------------------------------------
	// 创建按钮
	but_ts := widget.NewButton("json-gen-ts", func() {
		println("Button clicked!")
	})
	// 设置按钮的位置和大小
	but_ts.Resize(fyne.NewSize(100, 40)) // 宽高为 100x40
	but_ts.Move(fyne.NewPos(10, 10))     // 坐标 x=10, y=10
	//---------------------------------------------------------------------------
	// 创建按钮
	but_go := widget.NewButton("json-gen-go", func() {
		println("Button clicked!")
	})
	// 设置按钮的位置和大小
	but_go.Resize(fyne.NewSize(100, 40)) // 宽高为 100x40
	but_go.Move(fyne.NewPos(10, 60))     // 坐标 x=10, y=10
	//---------------------------------------------------------------------------
	// 创建按钮
	but_java := widget.NewButton("json-gen-java", func() {
		println("Button clicked!")
	})
	// 设置按钮的位置和大小
	but_java.Resize(fyne.NewSize(100, 40)) // 宽高为 100x40
	but_java.Move(fyne.NewPos(10, 110))    // 坐标 x=10, y=10

	// 使用自由布局
	myWindow.SetContent(container.NewWithoutLayout(
		but_ts,
		but_go,
		but_java,
	))

	// 显示窗口
	myWindow.ShowAndRun()
}
