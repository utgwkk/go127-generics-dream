package parser

// このコードはGo 1.27 (HEAD) でもコンパイルできません

// Parser は入力の先頭を消費して型 T の値を生成する。
type Parser[T any] struct {
	run func(input string) (value T, rest string, ok bool)
}

func (p Parser[T]) Parse(input string) (T, string, bool) {
	return p.run(input)
}

// Then の結果を保持するペア。
type Pair[A, B any] struct {
	First  A
	Second B
}

// Satisfy は述語を満たす1文字にマッチ（簡略化のためASCII前提）。
func Satisfy(pred func(rune) bool) Parser[rune] {
	return Parser[rune]{run: func(in string) (rune, string, bool) {
		if len(in) == 0 {
			return 0, in, false
		}
		r := rune(in[0])
		if !pred(r) {
			return 0, in, false
		}
		return r, in[1:], true
	}}
}

func Char(c rune) Parser[rune] {
	return Satisfy(func(r rune) bool { return r == c })
}

// Map: 結果型 T を U へ変換 ← メソッド型パラメータ U が必須
func (p Parser[T]) Map[U any](f func(T) U) Parser[U] {
	return Parser[U]{run: func(in string) (U, string, bool) {
		v, rest, ok := p.run(in)
		if !ok {
			var zero U
			return zero, in, false
		}
		return f(v), rest, true
	}}
}

// Then: p の後に q を実行し、結果をペアに ← 型が T から Pair[T,U] へ変化
func (p Parser[T]) Then[U any](q Parser[U]) Parser[Pair[T, U]] {
	return Parser[Pair[T, U]]{run: func(in string) (Pair[T, U], string, bool) {
		a, rest1, ok := p.run(in)
		if !ok {
			var zero Pair[T, U]
			return zero, in, false
		}
		b, rest2, ok := q.run(rest1)
		if !ok {
			var zero Pair[T, U]
			return zero, in, false
		}
		return Pair[T, U]{a, b}, rest2, true
	}}
}

// SkipThen: p を実行して結果を捨て、q の結果(型 U)を返す
func (p Parser[T]) SkipThen[U any](q Parser[U]) Parser[U] {
	return p.Then(q).Map(func(pr Pair[T, U]) U { return pr.Second })
}

// ThenSkip: p の後に q を実行し、p の結果(型 T)を保持
func (p Parser[T]) ThenSkip[U any](q Parser[U]) Parser[T] {
	return p.Then(q).Map(func(pr Pair[T, U]) T { return pr.First })
}
