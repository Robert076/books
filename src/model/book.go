package book

type Book struct {
	Id          int
	Title       string
	Author      string
	ReleaseYear string
	Topic       string
}

var Books = []Book{
	{1, "Clean Code", "Robert C. Martin", "2008", "Software Engineering"},
	{2, "The Pragmatic Programmer", "Andrew Hunt, David Thomas", "1999", "Software Development"},
	{3, "Introduction to Algorithms", "Thomas H. Cormen", "2009", "Algorithms"},
	{4, "Design Patterns", "Erich Gamma et al.", "1994", "Software Architecture"},
	{5, "You Don't Know JS", "Kyle Simpson", "2015", "JavaScript"},
	{6, "Eloquent JavaScript", "Marijn Haverbeke", "2018", "JavaScript"},
	{7, "Automate the Boring Stuff with Python", "Al Sweigart", "2015", "Python"},
	{8, "Code Complete", "Steve McConnell", "2004", "Software Engineering"},
	{9, "Structure and Interpretation of Computer Programs", "Harold Abelson, Gerald Jay Sussman", "1996", "Computer Science"},
	{10, "Python Crash Course", "Eric Matthes", "2016", "Python"},
	{11, "Refactoring", "Martin Fowler", "1999", "Software Engineering"},
	{12, "Head First Design Patterns", "Eric Freeman, Elisabeth Robson", "2004", "Software Design"},
	{13, "The Art of Computer Programming", "Donald E. Knuth", "1968", "Computer Science"},
	{14, "JavaScript: The Good Parts", "Douglas Crockford", "2008", "JavaScript"},
	{15, "Grokking Algorithms", "Aditya Bhargava", "2016", "Algorithms"},
}
