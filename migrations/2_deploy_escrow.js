// migrations/2_deploy_escrow.js

const Escrow = artifacts.require("Escrow");

module.exports = function (deployer) {
    const buyer = "0x48B1Fe8A1014701C854620bBF4498184a5610Bc2";
    const seller = "0x5726B14F9afCF854b1Cee9FD3b099C66E50faa78";
    const trustee = "0x0Cde631D6E0254eFE29f42b953EE6AF669e17D21";
    const amountINR = 10000;
    const buyerBank = {
        accountHolderName: "Buyer Name",
        accountNumber: "2323230085668044",
        bankName: "Buyer Bank",
        ifscCode: "IFSC0001"
    };
    const sellerBank = {
        accountHolderName: "Gaurav Kumar",
        accountNumber: "1121431121541121",
        bankName: "HDFC Bank",
        ifscCode: "HDFC0001234"
    };
    const payoutAuth = {
        razorpayApiKey: "rzp_test_aosIPPw8Fvw5EC",
        razorpayApiSecret: "4aQUVBgZZrtPIW6BcwaVBANL"
    };

    deployer.deploy(Escrow, buyer, seller, trustee, amountINR, buyerBank, sellerBank, payoutAuth);
};
