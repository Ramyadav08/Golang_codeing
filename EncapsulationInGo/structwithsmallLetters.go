package encapsulation


type Person struct {
	fname string
	lname string
	age int
}


// constructtor

func NewPerson(fn,ln string, a int) Person{
	return Person{
		fname:fn,
		lname:ln,
		age:a,
	}
}


func (p *Person)SetName(f string){
	p.fname=f
}

func (p *Person)GetName()string{
	return p.fname
}