package main

import "fmt"
import "math/rand"
import "time"
import "strconv"

func main() {
        const max int32 = 20
        var answer int64
        seed := time.Now().UnixNano()
        rnd := rand.New(rand.NewSource(seed))
        secret := int64(rnd.Int31n(max))
        fmt.Printf("Hi, i'm thinking of a number between 0 and %d. Can you guess it?\n", max)
        answer = getinput()
        for {
                if answer > secret {
                        fmt.Println("My number is smaller")
                }
                if answer < secret {
                        fmt.Println("My number is bigger")
                }
                if answer == secret {
                        fmt.Println("Well done, You guessed it!")
                        break
                }
                answer = getinput()
        }
}

func getinput() int64 {
        var stranswer string
        var number int64
        var err error
        fmt.Print("Guess my number: ")
        _, _ = fmt.Scanf("%s", &stranswer)
        number, err = strconv.ParseInt(stranswer, 10, 32)
        for err != nil {
                fmt.Println("Is that a number ?")
                fmt.Print("Guess my number: ")
                _, _ = fmt.Scanf("%s", &stranswer)
                number, err = strconv.ParseInt(stranswer, 10, 32)
        }
        return number
}
