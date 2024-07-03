// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Escrow {
    address public buyer;
    address public seller;
    address public trustee;
    uint256 public amountINR;

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
    event PayoutRequested(string buyerAccount, string sellerAccount, uint256 amountINR);

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
    }

    function approveByBuyer() public {
        require(msg.sender == buyer, "Only buyer can approve");
        buyerApproval = true;
        emit ApprovalReceived(msg.sender);
        checkAndRequestPayout();
    }

    function approveBySeller() public {
        require(msg.sender == seller, "Only seller can approve");
        sellerApproval = true;
        emit ApprovalReceived(msg.sender);
        checkAndRequestPayout();
    }

    function approveByTrustee() public {
        require(msg.sender == trustee, "Only trustee can approve");
        trusteeApproval = true;
        emit ApprovalReceived(msg.sender);
        checkAndRequestPayout();
    }

    function disapproveByBuyer() public {
        require(msg.sender == buyer, "Only buyer can disapprove");
        buyerApproval = false;
        emit ApprovalReceived(msg.sender);
    }

    function disapproveBySeller() public {
        require(msg.sender == seller, "Only seller can disapprove");
        sellerApproval = false;
        emit ApprovalReceived(msg.sender);
    }

    function checkAndRequestPayout() internal {
        if ((buyerApproval && sellerApproval) || (trusteeApproval && (buyerApproval || sellerApproval))) {
            emit PayoutRequested(buyerBankDetails.accountNumber, sellerBankDetails.accountNumber, amountINR);
        }
    }
}
