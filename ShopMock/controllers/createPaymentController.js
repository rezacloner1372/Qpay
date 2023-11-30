const moment = require('jalali-moment');
const axios = require('axios');
const qs = require('qs');
const Payment = require("../models/payment");
const User = require("../models/user");

const createPaymentController = async (req, res, next) => {
    try {
        const newUser = await new User({
            fullName: req.body.fullName,
            phone: req.body.phone,
        });

        let savedUser = await newUser.save();

        const newPayment = await new Payment({
            user: savedUser._id,
            amount: req.body.amount,
        });

        let savedPayment = await newPayment.save();

        const data = qs.stringify({
            Amount: savedPayment.amount,
            callback_url: `${process.env.QPAY_SHOP_MOCK_APP_EXTERNAL_URL}/api/v1/payment/verify`,
            Description: `نام و نام خانوادگی مشتری: ${req.body.fullName} | شماره تماس مشتری: ${req.body.phone}`,
            Phone: req.body.phone,
            Email: "test@test.com",
            merchant_id: process.env.QPAY_MERCHANT_ID,
        });
        let config = {
            method: 'post',
            maxBodyLength: Infinity,
            url: `${process.env.QPAY_API_URL}/payment/request`,
            headers: { 
              'Content-Type': 'application/x-www-form-urlencoded', 
            },
            data : data
        };

        const response = await axios.request(config);

        console.log(response.data);
        console.log(response.status);
        
        if (response.status != 200) {
            return res.status(400).send({
                status: 400,
                message: "",
                messageFA: "خطا در انتقال به درگاه پرداخت، لطفا دوباره تلاش کنید."
            });
        }

        let newQpayPayment = await Payment.findOneAndUpdate({
            _id: savedPayment._id
        }, {
            authority: response.data.authority,
        }, {
            new: true
        })

        return res.send({
            status: 200,
            message: "",
            messageFA: "در حال انتقال به درگاه پرداخت ...",
            redirectUrl: `${process.env.QPAY_API_URL}/payment/${newQpayPayment.authority}`,
        });

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
    createPaymentController
}