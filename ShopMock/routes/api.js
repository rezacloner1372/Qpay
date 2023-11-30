const express = require("express");
const router = express.Router();


const {
    createPaymentController
} = require('../controllers/createPaymentController');

const {
    verifyPaymentController
} = require('../controllers/verifyPaymentController');

const {
    getAllPaymentsController
} = require('../controllers/getAllPaymentsController');


router.post("/v1/payment/create", [], createPaymentController);
router.get("/v1/payment/verify", [], verifyPaymentController);
router.get("/v1/payment/all", [], getAllPaymentsController);


module.exports = router;