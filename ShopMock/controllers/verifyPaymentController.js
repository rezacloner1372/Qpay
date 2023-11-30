const moment = require('jalali-moment');
const axios = require('axios');
const qs = require('qs');

const Payment = require("../models/payment");
const User = require("../models/user");

const verifyPaymentController = async (req, res, next) => {
    try {
        if (!req.query.status || req.query.status != "successful") {
            console.log("status is not successful");
            return res.redirect(`${process.env.QPAY_SHOP_MOCK_APP_EXTERNAL_URL}/?paymentStatus=NOK`);
        }
        if (!req.query.authority) {
            console.log("authority is not provided");
            return res.redirect(`${process.env.QPAY_SHOP_MOCK_APP_EXTERNAL_URL}/?paymentStatus=NOK`);
        }

        let payment = await Payment.findOne({
            authority: req.query.authority
        })

        if (!payment) {
            console.log("payment not found");
            return res.redirect(`${process.env.QPAY_SHOP_MOCK_APP_EXTERNAL_URL}/?paymentStatus=NOK`);
        }

        const data = qs.stringify({
            Amount: payment.amount,
            authority: payment.authority,
            merchant_id: process.env.QPAY_MERCHANT_ID,
        });
        let config = {
            method: 'post',
            maxBodyLength: Infinity,
            url: `${process.env.QPAY_API_URL}/payment/verify`,
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            data: data
        };

        const response = await axios.request(config);

        console.log(response.data);

        if (response.status != 200) {
            console.log("payment verification failed");
            return res.redirect(`${process.env.QPAY_SHOP_MOCK_APP_EXTERNAL_URL}/?paymentStatus=NOK`);
        }

        await Payment.findOneAndUpdate({
            _id: payment._id
        }, {
            isPaymentSuccessful: true,
        }, {
            new: true
        })
        return res.redirect(`${process.env.QPAY_SHOP_MOCK_APP_EXTERNAL_URL}/?paymentStatus=OK`);
    } catch (err) {
        console.log(err);
        return res.status(500).send({
            status: 500,
            message: "Internal Server Error",
            messageFA: "خطای داخلی سرور ، چند ثانیه دیگر تلاش فرمایید."
        });
    }
}

module.exports = {
    verifyPaymentController
}