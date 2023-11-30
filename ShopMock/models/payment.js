const mongoose = require("mongoose");
const AutoIncrement = require('mongoose-sequence')(mongoose);

const Schema = mongoose.Schema;

const paymentSchema = Schema({
    user: {
        type: Schema.Types.ObjectId,
        ref: "User"
    },
    authority: {
        type: String,
        trim: true
    },
    isPaymentSuccessful: {
        type: Boolean,
        default: false
    },
    amount: {
        type: Number,
        required: true
    },
    createdAt: {
        type: Date,
        default: Date.now,
    }
});
paymentSchema.plugin(AutoIncrement, {
    inc_field: 'paymentId',
    start_seq: 1000000
});
module.exports = mongoose.model("Payment", paymentSchema);