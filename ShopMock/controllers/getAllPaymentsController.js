const Payment = require("../models/payment");
const User = require("../models/user");

const getAllPaymentsController = async (req, res) => {
    try {
        const payments = await Payment.find().populate("user");
        return res.send({
            status: 200,
            message: "",
            data: payments
        });
    } catch (err) {
        return res.status(500).send({
            status: 500,
            message: "",
            messageFA: "خطا در دریافت لیست پرداخت ها"
        });
    }
}

module.exports = {
    getAllPaymentsController
}