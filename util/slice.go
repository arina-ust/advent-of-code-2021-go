package util


func Flatten[T any](lists [][]T) []T {
    var res []T
    for _, list := range lists {
        res = append(res, list...)
    }
    return res
}

func RemoveWhiteSpaces(list []string) []string {
    var result []string
    for _, v := range list {
        if len(v) == 0 {
            continue
        }
        result = append(result, v)
    }
    return result
}