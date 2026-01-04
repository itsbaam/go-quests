# Go Structs Quest — Immutable User Profile

Your task is to implement a **user profile system** using Go structs and methods.

This quest focuses on:

* Defining and initializing structs
* Value vs pointer receivers
* Updating struct fields safely
* Returning copies vs mutating state
* Encapsulation through methods
* Avoiding unintended mutations

---

### Reference

* [https://gobyexample.com/structs](https://gobyexample.com/structs)
* [https://gobyexample.com/methods](https://gobyexample.com/methods)
* [https://go.dev/tour/moretypes/2](https://go.dev/tour/moretypes/2)

---

### Struct Definition

```go
type User struct {
	ID    int
	Name  string
	Age   int
	Email string
}
```

---

### Functions & Methods to Implement

```go
func NewUser(id int, name, email string, age int) *User

func (u *User) IsAdult() bool
func (u User) DisplayName() string

func (u *User) UpdateEmail(email string)
func (u *User) Birthday()

func CloneUser(u User) User
```

---

### Description

You are building a **simple user profile domain model**. The struct represents a user’s **state**, and methods define valid interactions with that state.

Some methods must **mutate the user**, while others must be **read-only**. Tests will ensure that you understand the difference.

---

### Function Behavior (Very Explicit)

---

#### 1. `NewUser`

Creates and initializes a new user.

```go
u := NewUser(1, "Mani", "mani@example.com", 20)
```

**Rules:**

* All fields must be assigned from parameters
* The returned value must not be `nil`
* The function returns a pointer to a `User`

---

#### 2. `IsAdult`

Checks whether the user is an adult.

```go
u := &User{Age: 18}
u.IsAdult() // true
```

**Rules:**

* Return `true` if `Age >= 18`
* Return `false` otherwise
* Must not modify the user

---

#### 3. `DisplayName`

Returns a formatted display string for the user.

```go
u := User{Name: "Mani", Email: "mani@example.com"}
u.DisplayName() // "Mani <mani@example.com>"
```

If fields are empty:

```go
u := User{}
u.DisplayName() // " <>"
```

**Rules:**

* Use a **value receiver**
* Must not modify the user

---

#### 4. `UpdateEmail`

Updates the user’s email address.

```go
u := &User{Email: "old@example.com"}
u.UpdateEmail("new@example.com")
```

**Rules:**

* Must mutate the original user
* Must use a **pointer receiver**

---

#### 5. `Birthday`

Increments the user’s age by one.

```go
u := &User{Age: 20}
u.Birthday()
u.Age // 21
```

**Rules:**

* Must mutate the original user

---

#### 6. `CloneUser`

Creates a copy of a user.

```go
u1 := User{Age: 20}
u2 := CloneUser(u1)

u2.Birthday()
u1.Age // still 20
```

**Rules:**

* Returned user must be an **independent copy**
* Changes to the clone must not affect the original

---

### Requirements

1. Use a struct, not a map
2. Use both value and pointer receivers correctly
3. Do not return pointers unless explicitly asked
4. No global state
5. No unnecessary allocations
6. Fields must be accessed through methods where applicable

---

### Common Beginner Mistakes (Tests Will Catch These)

* Using pointer receivers everywhere
* Mutating state in value-receiver methods
* Returning pointers incorrectly
* Forgetting to dereference in mutating methods
* Confusing copying with aliasing

---

### Key Ideas (What This Quest Teaches)

* Structs model **state + behavior**
* Value receivers protect immutability
* Pointer receivers enable mutation
* Copying structs is cheap and explicit
* Methods define a domain boundary

---

### Run Tests

```bash
go test ./quests/10.structs -v
```
