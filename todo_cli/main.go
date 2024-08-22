package main

func main() {
	todos := Todos{}
	todos.add("complete this project")
	todos.add("complete the minor project")
	todos.add("ye kaam kyu nhi kr rha h")

	todos.toggle(0)

	todos.print()

	// fmt.Printf("%+v\n\n", todos)

}
