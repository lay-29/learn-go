package No4

import "fmt"

type LogMsgHandle func(string) (error, string)
type LogMsgEvent []LogMsgHandle

func (logEvent *LogMsgEvent) Add(handler LogMsgHandle) {
	*logEvent = append(*logEvent, handler)
}
func (logEvent *LogMsgEvent) Fire(msg string) {
	for _, handler := range *logEvent {
		if err, _ := handler(msg); err != nil {
			println(err.Error())
		}
	}
}
func FuncTypeTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()

	var lgm LogMsgHandle = func(msg string) (error, string) {
		fmt.Println(msg)
		return nil, "Success"
	}
	err, msg := lgm("1 -- log message")
	fmt.Println(err, msg)

	var logEvent *LogMsgEvent = nil
	logEvent.Add(func(msg string) (error, string) {
		fmt.Println("warning：", msg)
		return nil, "Warning"
	})
	logEvent.Add(func(msg string) (error, string) {
		fmt.Println("Info：", msg)
		return nil, "Info"
	})
	logEvent.Fire("2 -- log message")
}
