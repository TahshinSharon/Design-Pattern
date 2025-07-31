// Single Responsibility -> Each class / method should have only one responsibility

import java.io.FileWriter;
import java.io.IOException;

class BankAccount {
    private String accountNumber;
    private String accountHolder;
    private double balance;

    public BankAccount(String accountNumber, String accountHolder) {
        this.accountNumber = accountNumber;
        this.accountHolder = accountHolder;
        this.balance = 0.0;
    }

    public void deposit(double amount) {
        if (amount > 0) {
            balance += amount;
        }
    }

    public void withdraw(double amount) {
        if (amount > 0 && balance >= amount) {
            balance -= amount;
        }
    }

    public String getAccountNumber() {
        return accountNumber;
    }

    public String getAccountHolder() {
        return accountHolder;
    }

    public double getBalance() {
        return balance;
    }
}

class AccountPrinter {
    public void printDetails(BankAccount account) {
        System.out.println("Account Holder: " + account.getAccountHolder());
        System.out.println("Account Number: " + account.getAccountNumber());
        System.out.println("Balance: $" + account.getBalance());
    }
}

class AccountSaver {
    public void saveToFile(BankAccount account, String filename) {
        try (FileWriter writer = new FileWriter(filename)) {
            writer.write("Account Holder: " + account.getAccountHolder());
            writer.write("Account Number: " + account.getAccountNumber());
            writer.write("Balance: $" + account.getBalance());
        } catch (IOException e) {
            System.out.println("Error saving account: " + e.getMessage());
        }
    }
}

public class srp {
    public static void main(String[] args) {
        BankAccount account = new BankAccount("123456", "Tahshin");
        account.deposit(2000);
        account.withdraw(500);

        AccountPrinter printer = new AccountPrinter();
        printer.printDetails(account);

        AccountSaver saver = new AccountSaver();
        saver.saveToFile(account, "Tahshin_Bank_Details.txt");
    }
}
