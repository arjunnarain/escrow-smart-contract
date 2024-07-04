// migrations/2_deploy_escrow.js

const Escrow = artifacts.require("Escrow");

module.exports = function (deployer) {
    const buyer = "0xe1621A94B849b994D9317635682B017D7aeAcEfb";
    const seller = "0x4a821D1155802cEF86C5d93445caa10f16D35569";
    const trustee = "0x66E40344870447841Ea02C35596F1C5669F08c3A";
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
