func myFunc(a []int) {
    if len(a) == 0 {
        fmt.Println("Empty slice!")
        return
    }
    for i := range a {
        if a[i] == 0 {
            fmt.Println("Found zero!")
            return
        }
    }
}