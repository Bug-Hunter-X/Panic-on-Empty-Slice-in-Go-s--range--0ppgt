func myFunc(a []int) {
    for i := range a {
        if a[i] == 0 {
            fmt.Println("Found zero!")
            return
        }
    }
}