package main

import(
  "fmt"
)

func earningBeforeTaxCalculator(revenue float64 , expenses float64) float64 {
   return revenue-expenses;
}

func earningAfterTaxCalculator(revenue, expenses, taxRate float64) float64 {
   taxAmount := revenue*taxRate/100;
   return revenue-(expenses+taxAmount);
}

func profitCalculator(){

      var revenue,expenses,taxRate float64;

      fmt.Print("Enter Revenue: ");
	  fmt.Scan(&revenue);

	  fmt.Print("Enter Expenses: ");
	  fmt.Scan(&expenses);


	  fmt.Print("Enter Tax Rate: ");
      fmt.Scan(&taxRate);

	  earningBeforeTaxAmount := earningBeforeTaxCalculator(revenue, expenses);
	  earningAfterTaxAmount  := earningAfterTaxCalculator(revenue, expenses, taxRate);
      ratio := earningAfterTaxAmount/earningBeforeTaxAmount;
      message := fmt.Sprintf("Earning befor tax amount is %v and earnings after tax amount is %v and ratio is %v", earningBeforeTaxAmount,earningAfterTaxAmount,ratio); 
      fmt.Println(message);


	//   fmt.Println("Earning Before Tax Amount: ", earningBeforeTaxAmount);
	//   fmt.Println("Earning After Tax Amount: ", earningAfterTaxAmount);
	//   fmt.Println("Ratio: ", earningAfterTaxAmount / earningBeforeTaxAmount);


}


func getUserInput(infoText string) float64{
	var userInput float64;
	fmt.Print(infoText);
	fmt.Scan(&userInput);
	return userInput;
}


func profitCalculatorV2(){
  fmt.Println("v2 is starting....");
  revenue  :=  getUserInput("Enter Revenue: ");
  expenses :=  getUserInput("Enter Expenses: ");
  taxRate  :=  getUserInput("Enter Tax Rate: ");

	earningBeforeTaxAmount, earningAfterTaxAmount, ratio := profitCalculatorHelper(revenue, expenses, taxRate);
	message := fmt.Sprintf("Earning befor tax amount is %v and earnings after tax amount is %v and ratio is %v", earningBeforeTaxAmount,earningAfterTaxAmount,ratio); 
	fmt.Println(message);

}

func profitCalculatorHelper(revenue, expenses, taxRate float64) (float64,float64,float64) {
	earningBeforeTaxAmount := earningBeforeTaxCalculator(revenue, expenses)
	earningAfterTaxAmount  := earningAfterTaxCalculator(revenue, expenses, taxRate)
	ratio := earningAfterTaxAmount/earningBeforeTaxAmount
	return earningBeforeTaxAmount, earningAfterTaxAmount, ratio
}