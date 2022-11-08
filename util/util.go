package util

func Strptr(c string) *string {
    var v string = c
    return &v
}

func Strptrs(cs ...string) []*string {
    var strs []*string
    for _, c := range cs {
        var t string = c
        strs = append(strs, &t)
    }
    return strs
}

func Derefstr(sp []*string) []string {
    var str []string
    for _, s := range sp {
        str = append(str, *s)
    }
    return str
}