// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Escrow {
    address public buyer;
    address public seller;
    address public trustee;
    uint256 public amountINR;

    enum State { Active, Completed }
    State public currentState;

    struct BankDetails {
        string accountHolderName;
        string accountNumber;
        string bankName;
        string ifscCode;
    }

    BankDetails public buyerBankDetails;
    BankDetails public sellerBankDetails;

    struct PayoutAuth {
        string razorpayApiKey;
        string razorpayApiSecret;
    }

    PayoutAuth public payoutAuth;

    bool public buyerApproval;
    bool public sellerApproval;
    bool public trusteeApproval;

    event ApprovalReceived(address indexed approver);
    event PayoutRequested(
        string buyerAccountHolderName,
        string buyerAccountNumber,
        string buyerBankName,
        string buyerIfscCode,
        string sellerAccountHolderName,
        string sellerAccountNumber,
        string sellerBankName,
        string sellerIfscCode,
        uint256 amountINR
    );
    event ContractCompleted();

    modifier onlyActive() {
        require(currentState == State.Active, "Contract is no longer active");
        _;
    }

    modifier canComplete() {
        require(
            msg.sender == trustee,
            "Only trustee can complete the contract"
        );
        _;
    }

    event EscrowCreated(
        address indexed buyer,
        address indexed seller,
        address indexed trustee,
        uint256 amountINR
    );

    constructor(
        address _buyer,
        address _seller,
        address _trustee,
        uint256 _amountINR,
        BankDetails memory _buyerBankDetails,
        BankDetails memory _sellerBankDetails,
        PayoutAuth memory _payoutAuth
    ) {
        buyer = _buyer;
        seller = _seller;
        trustee = _trustee;
        amountINR = _amountINR;
        buyerBankDetails = _buyerBankDetails;
        sellerBankDetails = _sellerBankDetails;
        payoutAuth = _payoutAuth;
        currentState = State.Active;

        emit EscrowCreated(_buyer, _seller, _trustee, _amountINR);
    }

    function approveByBuyer() public onlyActive {
        require(msg.sender == buyer, "Only buyer can approve");
        buyerApproval = true;
        emit ApprovalReceived(msg.sender);
        checkAndRequestPayout();
    }

    function approveBySeller() public onlyActive {
        require(msg.sender == seller, "Only seller can approve");
        sellerApproval = true;
        emit ApprovalReceived(msg.sender);
        checkAndRequestPayout();
    }

    function approveByTrustee() public onlyActive {
        require(msg.sender == trustee, "Only trustee can approve");
        trusteeApproval = true;
        emit ApprovalReceived(msg.sender);
        checkAndRequestPayout();
    }

    function disapproveByBuyer() public onlyActive {
        require(msg.sender == buyer, "Only buyer can disapprove");
        buyerApproval = false;
        emit ApprovalReceived(msg.sender);
    }

    function disapproveBySeller() public onlyActive {
        require(msg.sender == seller, "Only seller can disapprove");
        sellerApproval = false;
        emit ApprovalReceived(msg.sender);
    }

    function checkAndRequestPayout() internal {
        if ((buyerApproval && sellerApproval) || (trusteeApproval && (buyerApproval || sellerApproval))) {
            emit PayoutRequested(
                buyerBankDetails.accountHolderName,
                buyerBankDetails.accountNumber,
                buyerBankDetails.bankName,
                buyerBankDetails.ifscCode,
                sellerBankDetails.accountHolderName,
                sellerBankDetails.accountNumber,
                sellerBankDetails.bankName,
                sellerBankDetails.ifscCode,
                amountINR
            );
        }
    }

    function completeContract() public onlyActive canComplete {
        currentState = State.Completed;
        emit ContractCompleted();
    }
}