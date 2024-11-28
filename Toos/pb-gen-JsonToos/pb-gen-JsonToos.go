package main

import (
	"PbJsonToTsAndGo/Toos/json-gen-ts/to_ts"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
	"path/filepath"
	"time"
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
	case 1:
		to_ts.NewPbToTsMain().Start(dirPath)
	case 2:
	case 3:

	}
	m := NewPbToMain()
	// 检查文件夹是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 如果不存在则创建文件夹
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println("Failed to create directory:", err)
			return
		}
		fmt.Println("Directory created successfully!")
	}

	files, err := os.Open(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer files.Close()

	fileNames, err := files.Readdirnames(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	//go build excel_to_cx.go
	/*
		fmt.Printf("----------*.xlsx生成规则--------------------\n")
		fmt.Printf("--第一行 字段属性为 默认空:客户端和服务端使用;c:只有客户端用;s:只有服务端使用;\n")
		fmt.Printf("--第二行 字段属性名字 如果为空字符串这个列将不生成配置\n")
		fmt.Printf("--第三行 字段类型 如果为空 将默认为string,其它类型有 string,float\n")
		fmt.Printf("--第四行 字段介绍名称\n")
		fmt.Printf("--第五行 字段属性使用介绍\n")
		fmt.Printf("--第六行 配置第一行数据开始\n")
	*/
	fmt.Printf("added by yh @ 2024/11/27 17:35 408309839@qq.com \n")
	fmt.Printf("\n")
	if m.CheckToTime() {
		fmt.Printf("--少年有报错联系管理员...\n")
		m.Hang()
		return
	}
	go_output := "./Protocal" //"./bin/cfg_go"
	//json_output := "./cfg_json" //"./bin/cfg_json"
	m.DeleteFiles(go_output, ".ts")
	//m.DeleteFiles(json_output, ".json")

	for _, fileName := range fileNames {
		if filepath.Ext(fileName) == ".proto" {
			filePath := filepath.Join(dirPath, fileName)
			t := to_ts.NewPbToTsItem()
			t.OpenProtoFile(fileName, filePath, go_output)

			//js := to_json.NewExcelToJson()
			//js.OpenExcelFile(fileName, filePath, json_output)
		}
	}
	fmt.Printf("\n")
	fmt.Printf("----------生成所有C#配置完成---------------------\n")
	time.Sleep(2 * time.Second)

}
