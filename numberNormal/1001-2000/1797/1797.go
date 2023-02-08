package leetcode

type AuthenticationManager struct {
	timeToLive int
	record     map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{timeToLive, make(map[string]int)}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.record[tokenId] = currentTime
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if _, isPresent := this.record[tokenId]; isPresent {
		if this.record[tokenId]+this.timeToLive > currentTime {
			this.record[tokenId] = currentTime
		} else {
			delete(this.record, tokenId)
		}
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	res := 0

	for id, time := range this.record {
		if time+this.timeToLive > currentTime {
			res++
		} else {
			delete(this.record, id)
		}
	}

	return res
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
