package main

func main() {
	funWithPointers()
}

func funWithPointers() {
	i := 1
	print("i=")
	println(i)
	print("&i=")
	println(&i)

	p := &i
	print("p=")
	println(p)

	*p++
	println("p++")
	print("i=")
	println(i)
	print("p=")
	println(p)
}
