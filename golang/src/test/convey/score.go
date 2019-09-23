package main


func main(){
	ShowYouScore(100)
}

func ShowYouScore(score int) string {
     if score == 100 {
     	return "very good"
	 }

	 if score >= 90 {
	 	if score >=95 {
			return "good+"
		}
        return "good"
	 }

	 if score >= 80{
	 	return "medium"
	 }

	 if score >= 60 {
	 	 return "pass"
	 }

	 return "poor"
}
