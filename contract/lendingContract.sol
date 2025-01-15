// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

contract LendingContract {

    address public owner;
    uint256 public interestRate;
    uint256 public loanAmountLimit;

    struct Loan {
        uint256 amount;
        uint256 interest;
        uint256 dueDate;
        bool isPaid;
    }

    mapping(address => Loan) public loans;

    // LOG ACTION
    event LoanIssued(address borrower, uint256 amount, uint256 interest, uint256 dueDate);
    event LoanRepaid(address borrower, uint256 amountPaid);
    event LoanDefaulted(address borrower);

    // ACCESS DECLARE
    modifier onlyOwner() {
        // make sure only owner can access
        require(msg.sender == owner, "Only the Owner can perform this action");
        _;
    }

    modifier loanExists(address borrower) {
        require(loans[borrower].amount > 0, "No active loan for this address");
        _;
    }

    modifier loanNotPaid(address borrower) {
        require(!loans[borrower].isPaid, "");
        _;
    }

    constructor (uint256 _interestRate, uint256 _loanAmountLimit) payable {
        require(_interestRate > 0 && _interestRate <= 10000, "Interest rate must be between 1 and 10000 basis points");
        require(_loanAmountLimit > 0, "Loan amount limit must be greater than 0");

        owner = msg.sender;
        interestRate = _interestRate;
        loanAmountLimit = _loanAmountLimit;
    }

    function issueLoan(address borrower, uint256 amount) external onlyOwner {
        require(amount <= loanAmountLimit, "Loan amount exceeds limit");

        uint256 interest = (amount * interestRate) / 100;
        uint dueDate = block.timestamp + 30 days;

        loans[borrower] = Loan({
            amount: amount,
            interest: interest,
            dueDate: dueDate,
            isPaid: false
        });

        // triggered event declare
        emit LoanIssued(borrower, amount, interest, dueDate);
    }

    // func pay loan can access for address exist with send
    function repayLoan() external payable loanExists(msg.sender) loanNotPaid(msg.sender) {
        Loan storage loan = loans[msg.sender];
        uint256 totalAmount = loan.amount + loan.interest;

        require(msg.value >= totalAmount, "Insufficient funds to repay the loan");

        loan.isPaid = true;
        payable(owner).transfer(msg.value);

        emit LoanRepaid(msg.sender, msg.value);
    }


    function checkLoanStatus(address borrower) external view returns (uint256, uint256, uint256, bool) {
        Loan memory loan = loans[borrower];
        return (loan.amount, loan.interest, loan.dueDate, loan.isPaid);
    }

    function checkOverdueLoan(address borrower) external view returns (bool) {
        Loan memory loan = loans[borrower];
        if (block.timestamp > loan.dueDate && !loan.isPaid) {
            return true;
        }
        return false;
    }

    function updateInterestRate(uint256 newInterestRate) external onlyOwner {
        interestRate = newInterestRate;
    }

    function updateLoanAmountLimit(uint256 newLoanLimit) external onlyOwner {
        loanAmountLimit = newLoanLimit;
    }

    function withdraw() external onlyOwner {
        payable(owner).transfer(address(this).balance);
    }


}
