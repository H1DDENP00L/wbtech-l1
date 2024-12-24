package main

import "fmt"

/*
	Реализовать паттерн «адаптер» на любом примере.
*/

// Структура представляющая китайца
type chinesePerson struct {
	name string
}

// Интерфейс, описывающий того кто говорит по-китайски (Target interface)
type canSpeakChinese interface {
	speakChinese()
}

// Метод speakChinese() для структуры chinesePerson
func (cp chinesePerson) speakChinese() {
	fmt.Printf("我是%s 我会说中文!", cp.name)
}

// Структура представляющая русского
type russianPerson struct {
	name string
}

// Интерфейс, описывающий того кто говорит по-русски (Adaptee interface)
type canSpeakRussian interface {
	speakRussian()
}

// Метод speakRussian() для структуры russianPerson
func (rp russianPerson) speakRussian() {
	fmt.Printf("Я %s, и я говорю по-русски, но не знаю китайского!\n", rp.name)
}

// Adapter - структура для перевода
type translator struct {
	russianPerson
}

// реализация Target interface, через вызов Adaptee метода
func (t translator) speakChinese() {
	fmt.Printf("Умеет переводить с русского на китайский!!!. Перевел с русского фразу:\n")
	fmt.Println("Получилось - 我是弗拉基米尔，我说俄语，但我不懂中文!")
}

func main() {
	chineseGuy := chinesePerson{"李华"}
	russianGuy := russianPerson{"Владимир"}
	translateGuy := translator{russianPerson: russianGuy}

	fmt.Println("Chinese person:", chineseGuy)
	chineseGuy.speakChinese()
	fmt.Println("\n")

	fmt.Println("Russian person:", russianGuy)
	russianGuy.speakRussian()
	fmt.Println("\n")

	fmt.Println("Adapter:", translateGuy)
	translateGuy.speakChinese()
}
