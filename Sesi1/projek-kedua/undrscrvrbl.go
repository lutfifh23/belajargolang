// penggunaan garis bawah apabila ada variabel yang tidak dipakai
package main

func main() {
	var firstvariable string
	var name, age, address = "Aurell", 21, "Jakarta"

	_, _, _, _ = firstvariable, name, age, address
}
