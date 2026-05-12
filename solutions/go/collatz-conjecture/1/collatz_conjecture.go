package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
        return 0, errors.New("value cannot be negative")
    }

    if n == 1 {
        return 0, nil
    }
    
    steps := 0
	number := n
    
    for number != 1 {
        if number % 2 == 0 {
            number = number / 2
            steps++
        } else {
            number = (number * 3) + 1
            steps++
        }
    }

    return steps, nil
}
