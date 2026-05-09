package main

import ("fmt"; "math")

func main() {
  bannerWidth := 12
  fmt.Println("BannerWidth:", bannerWidth )

  bannerHeight := 8
  fmt.Println("BannerHeight:", bannerHeight)
  
  
  bannerArea := bannerWidth * bannerHeight
  fmt.Println("BannerArea:", bannerArea)

  halfBannerArea := bannerArea/2
  fmt.Println("HalfBannerArea:", halfBannerArea)
  
  bannerBorderLength := (bannerWidth + bannerHeight) * 2
  fmt.Println("BannerBorderLength:", bannerBorderLength)
  
  boxCount := 29
  Couriers := 5
  leftoverBoxes := boxCount % Couriers
  fmt.Println("LeftOverBoxes", leftoverBoxes)

  tempMorning := 18
  tempAfternoon := 30
  tempEvening := 23
  totalTemp := tempAfternoon + tempEvening + tempMorning
  fmt.Println("TotalTemp", totalTemp)
  averageTemp := totalTemp / 3
  fmt.Println("AverageTemp", averageTemp)

  knownWords := 47
  wordsGoal := 120
  progressPercent := (float64(knownWords)/float64(wordsGoal)) * 100
  fmt.Println("ProgressPercent: %.2f%%\n", progressPercent)

  coins := 0
  coins += 500
  coins += 1200
  coins /= 2
  coins -= 300
  fmt.Println("Coins:", coins)

  participants := 42
  groupCount := 8
  participantsPerGroup := participants/groupCount 
  fmt.Println("ParticipantsPerGroup:", participantsPerGroup)

  fmt.Println(20 - 4 * 3)
  fmt.Println((20 - 4) * 3)
   // потому что первое действие происходит в скобках
  
  squareValue := 81.0
  result := math.Sqrt(squareValue)
  fmt.Println("Result", result)
  multiplier := 5.0
  exponent := 2.0
  Pow := math.Pow(multiplier, exponent)
  fmt.Println("Pow", Pow)

}
