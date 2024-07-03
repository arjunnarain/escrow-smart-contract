// migrations/2_deploy_escrow.js

const Escrow = artifacts.require("Escrow");

module.exports = function (deployer) {
    const buyer = "0x2e0DdfC6Cd122571729b2B4063abF9320cdf112e";
    const seller = "0xB11d4E637cC1D9bB4ec3b1C8F828CA5Ee8574dce";
    const trustee = "0x913870a5C637F815dEf4dFC3eD38ed0667Db62bE";
    const amountINR = web3.utils.toWei("1", "ether"); // Adjust the amount as needed
    const buyerBank = {
        accountHolderName: "Buyer Name",
        accountNumber: "1234567890",
        bankName: "Buyer Bank",
        ifscCode: "IFSC0001"
    };
    const sellerBank = {
        accountHolderName: "Seller Name",
        accountNumber: "0987654321",
        bankName: "Seller Bank",
        ifscCode: "IFSC0002"
    };
    const payoutAuth = {
        razorpayApiKey: "your_razorpay_api_key",
        razorpayApiSecret: "your_razorpay_api_secret"
    };

    deployer.deploy(Escrow, buyer, seller, trustee, amountINR, buyerBank, sellerBank, payoutAuth);
};
