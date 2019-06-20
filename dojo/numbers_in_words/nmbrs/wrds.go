package nmbrs

// import "strconv"

func toTeen(input string) string {
  output := ""

  if input[len(input)-1:len(input)] == "y" {
    output = input[0:len(input)-1] + "een"
  } else {
    output = input + "teen"
  }

  return output
}

func parseInputToDecimal(input string) string {
  inputToDecimal := input[0:len(input)-3]
  return inputToDecimal + ".00 \\$"
}

func toDollars(size int) string {
  output := "dollar"
  if size != 1 {
    output = output + "s"
  }
  return output
}

func parseInputToLiteral(input string) string {
  unitNumbers := map[string]string{"1": "one", "2": "two", "3": "three", "4": "four", "5": "five", "6": "six", "7": "seven", "8": "eight", "9": "nine"}
  tenToTwelve := map[string]string{"10": "ten", "11": "eleven","12": "twelve"}
  // decimalNumbers := map[string]string{"two": "twenty", "three": "thirty", "four": "fourty", "five": "fifty", "six": "sixty", "seven": "seventy", "eight": "eighty", "nine": "ninety"}
  inputToNumber := input[0:len(input)-3]
  inputLength := len(inputToNumber)

  literalValue := ""
  switch inputLength {
    case 1: literalValue = unitNumbers[inputToNumber]
    case 2:
      if v, found := tenToTwelve[inputToNumber]; found {
        literalValue = v
      } else {
        // CONTINUE HERE
      }
  }

  return literalValue + " " + toDollars(inputLength)
}

func Wrds(input string) [2]string {
  outputValues := [2]string{parseInputToDecimal(input), "seven hundred and fourty five dollars"}

  return outputValues
}
