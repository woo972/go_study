package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

func (e Employee) calc(bonus int) int {
	return e.salary + bonus
}

func (e Employee) calcAnnual() int {
	return e.salary * 12
}

type Executive struct {
	Employee     // is a 관계  임원은 직원이다
	specialBonus int
}

func (e Executive) calcAnnual() int {
	return (e.salary * 12) + e.specialBonus
}

func main() {

	e1 := Employee{
		"kim",
		100,
	}
	fmt.Println("e1:", e1)
	fmt.Println(e1.calc(100) == 200)

	e2 := Executive{
		Employee: Employee{
			"lee",
			200,
		},
		specialBonus: 100,
	}
	fmt.Println("e2:", e2)
	fmt.Println(e2.calc(100) == 300)          // 임베디드 구조체의 매서드에 상속받은 것 처럼 접근할 수 있다.
	fmt.Println(e2.Employee.calc(100) == 300) // 하위로 찾아들어가는 것도 가능하다

	fmt.Println(e1.calcAnnual() == 1200)
	fmt.Println(e2.calcAnnual() == 2500) // 메서드가 오버라이딩된 것 처럼, 이름이 같아도 자신의 리시버에 먼저 접근한다.
	fmt.Println(e2.Employee.calcAnnual() == 2400)

}
