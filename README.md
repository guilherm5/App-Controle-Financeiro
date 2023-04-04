![Exemplo de uso](http://g.recordit.co/kerLONSUni.gif)

# Simple project for financial control

<p>This project was born out of a personal need to control my finances.</p>

## Functionality: 
- [x] Login user

- [x] Create salary

- [x] Create expenses

- [x] List salary

- [x] List expenses

- [x] Update salary

- [x] Update expenses

- [x] Delete salary

- [x] Delete salary

<p>This project has a login system, where with each login made, we will list the salaries and expenses of the logged-in user only. We also have a function that brings us the total amount spent on expenses, a function to subtract the value of Salary - Expenses etc..</p>

<p>This project separates expenses and salaries at each login, i.e. if you log in as user 1, you will only see user 1's salaries, expenses, and total expenses.</p>

## API made using
> Golang (Programming language)

> Gin (Framework Web)

> LibPQ (To handle postgresql)

> JWT (Login)

> Postgresql 

<p>We also haave the user verification, if the user does not exist in our database, it returns an error saying that the user does not exist.</p>

<p>
We also have jwt token verification, I did this using a middleware, if the token is invalid or expired, our routes will not be authenticated.</p>

![Exemplo de uso](http://g.recordit.co/kerLONSUni.gif)


