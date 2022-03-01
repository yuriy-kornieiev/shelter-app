package main

func main() {

	a := App{}
	a.Initialize("test", "test", "localhost", "3306", "raid_alert",
		"localhost", "6379", "")
	a.Run(":8080")

}
