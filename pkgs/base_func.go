package pkgs

/*
Swap 交换任意两个数值
a,b = b,a 并不能正确的进行交换, 是因为go默认是值传递，只有使用*的时候，才可以间接使用引用传递
*/
func Swap[T any](a, b *T) {
	//a, b = b, a
	*a, *b = *b, *a
}
