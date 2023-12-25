package main 

import(
	"fmt"
)
const amountBalance float64 = 1300000;

func printText(text string) {
	fmt.Println(text);
}

func welcomeToBank(){
	printText("Welcome to Go Bank!")
	printText("What do you want to do? ");
	printText("1. Check Balance");
	printText("2. Deposit");
	printText("3. Withdraw Money");
	printText("4. Exit");

}

func enterChoice() int {
	var choice int;
	fmt.Print("Enter your choice...:")
	fmt.Scan(&choice);
	return choice;
}

func checkBalance() string {
	currentBalance := fmt.Sprintf("Your balance is %.4f \n",amountBalance);
	return currentBalance;
}

func deposit() string {
     
	fmt.Print("Enter the amount you want to  deposit: ");
	var amountDeposited float64;
	fmt.Scan(&amountDeposited);
	amountDeposited = amountDeposited + amountBalance;
	return fmt.Sprintf("\nYour new balance is %.3f \n",amountDeposited);

}

func withdrawMoney() string {
      
	fmt.Print("Enter the amount to withdraw: ");
	var amountWithdrawn float64;
	fmt.Scan(&amountWithdrawn);
	amountWithdrawn = amountBalance - amountWithdrawn;
	return fmt.Sprintf("\nYour new balance is %.3f \n",amountWithdrawn);

}

func BankOperation (choice int) string  {

	if choice == 1 {
		return checkBalance();
	} else if choice == 2 {
		return deposit();
	} else if choice == 3 {
        return withdrawMoney();
	} else {
		return "Goodbye"
	}

}

func bank(){
	
    welcomeToBank()
	choice := enterChoice()
	message := BankOperation(choice)
    fmt.Println(message);

}