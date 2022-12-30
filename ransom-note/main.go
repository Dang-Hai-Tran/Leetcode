package ransomnote


func canConstruct(ransomNode string, magazine string) bool {
	mapMagazine := make(map[rune]int)
	mapRansomNode := make(map[rune]int)
	for _, c := range magazine {
		mapMagazine[c] += 1
	}
	for _, c := range ransomNode {
		mapRansomNode[c] += 1
	}
	for _, c := range ransomNode {
		if mapMagazine[c] < mapRansomNode[c] {
			return false
		}
	}
	return true
}
