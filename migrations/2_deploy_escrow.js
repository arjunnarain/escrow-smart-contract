// migrations/2_deploy_escrow.js

const Escrow = artifacts.require("Escrow");

module.exports = function (deployer) {
    const buyer = "0x62CA0844DA21BA467Ba11892f1859Fa939aeAa51";
    const seller = "0xD192d17d8bB82C80477683A0A865d28d33E0DE6E";
    const trustee = "0x09aC00C740bc2AF67c9d2e5e240e1154870a1330";
    const amountINR = 200;
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
